package app

import (
	"context"
	"sync"
	"time"

	"github.com/balabanovds/smonitor/internal/models"

	"github.com/balabanovds/smonitor/internal/collector"

	"github.com/balabanovds/smonitor/internal/collector/parsers"
	"github.com/balabanovds/smonitor/internal/metrics"
)

type App struct {
	storage     metrics.Storage
	parsers     []func() parsers.Parser
	parserTypes []models.ParserType
	interval    time.Duration
	timeout     time.Duration
	deleteOld   time.Duration
}

type InMetricChan <-chan models.Metric

func New(cfg Config, storage metrics.Storage) *App {
	var parserFuncs []func() parsers.Parser

	for _, p := range cfg.Parsers {
		pr := getParserFunc(p)
		if pr != nil {
			parserFuncs = append(parserFuncs, pr)
		}
	}

	return &App{
		storage:     storage,
		parsers:     parserFuncs,
		parserTypes: cfg.Parsers,
		interval:    time.Duration(cfg.Interval) * time.Second,
		timeout:     time.Duration(cfg.Timeout) * time.Second,
		deleteOld:   time.Duration(cfg.DeleteOld) * time.Second,
	}
}

func (a *App) Run(ctx context.Context) <-chan struct{} {
	doneCh := make(chan struct{})
	go func() {
		tick := time.NewTicker(a.interval)
		defer func() {
			tick.Stop()
			close(doneCh)
		}()

		for {
			select {
			case <-ctx.Done():
				return
			case t := <-tick.C:
				a.storage.Delete(t.Add(-a.deleteOld))
				select {
				case <-ctx.Done():
					return
				default:
				}

				for m := range a.worker(ctx) {
					a.storage.Save(m)
				}
			}
		}
	}()
	return doneCh
}

// fan-out - fan-in pattern.
func (a *App) worker(ctx context.Context) InMetricChan {
	streams := make([]InMetricChan, len(a.parsers))
	for i, parFn := range a.parsers {
		p := parFn()
		select {
		case <-ctx.Done():
			break
		default:
		}
		streams[i] = shifter(ctx, p.Parse(ctx))
	}
	return joinChannels(ctx, streams...)
}

func shifter(ctx context.Context, inCh <-chan collector.Result) InMetricChan {
	outCh := make(chan models.Metric)
	go func() {
		defer close(outCh)
		for {
			select {
			case <-ctx.Done():
				return
			case r, ok := <-inCh:
				if !ok {
					return
				}
				if r.Err != nil {
					// TODO log here
					continue
				}
				select {
				case <-ctx.Done():
					return
				case outCh <- r.Data:
				}
			}
		}
	}()
	return outCh
}

func joinChannels(ctx context.Context, streams ...InMetricChan) InMetricChan {
	var wg sync.WaitGroup
	outCh := make(chan models.Metric)

	multiplex := func(ch InMetricChan) {
		defer wg.Done()
		for i := range ch {
			select {
			case <-ctx.Done():
				return
			case outCh <- i:
			}
		}
	}

	wg.Add(len(streams))
	for _, st := range streams {
		go multiplex(st)
	}

	go func() {
		wg.Wait()
		close(outCh)
	}()

	return outCh
}

func getParserFunc(pType models.ParserType) func() parsers.Parser {
	switch pType {
	case models.LoadAvg:
		return parsers.NewLoadAvgParser
	case models.CPU:
		return parsers.NewCPUParser
	case models.IO:
		return parsers.NewIOParser
	case models.FS:
		return parsers.NewFSParser
	case models.Net:
		return parsers.NewNETParser
	case models.Undef:
	}
	return nil
}
