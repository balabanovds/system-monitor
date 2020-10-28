package app

import (
	"context"
	"sync"
	"time"

	"github.com/balabanovds/system-monitor/cmd/config"
	"github.com/balabanovds/system-monitor/internal/collector"
	"github.com/balabanovds/system-monitor/internal/collector/parsers"
	"github.com/balabanovds/system-monitor/internal/models"
	"github.com/balabanovds/system-monitor/internal/storage"
	"go.uber.org/zap"
)

type App struct {
	storage                storage.Storage
	parsers                []func() parsers.Parser
	parserTypes            []models.ParserType
	interval               time.Duration
	MaxMeasurementDuration time.Duration
	log                    *zap.Logger
}

type InMetricChan <-chan models.Metric

func New(cfg config.AppConfig, storage storage.Storage, logger *zap.Logger) *App {
	var parserFuncs []func() parsers.Parser

	for _, p := range cfg.Parsers {
		pr := getParserFunc(p)
		if pr != nil {
			parserFuncs = append(parserFuncs, pr)
		}
	}

	return &App{
		storage:                storage,
		parsers:                parserFuncs,
		parserTypes:            cfg.Parsers,
		interval:               time.Duration(cfg.IntervalSeconds) * time.Second,
		MaxMeasurementDuration: time.Duration(cfg.MaxMeasurementHours) * time.Hour,
		log:                    logger,
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
				a.storage.Delete(t.Add(-a.MaxMeasurementDuration))
				select {
				case <-ctx.Done():
					return
				default:
				}

				for m := range a.createChan(ctx) {
					a.storage.Save(m)
				}
			}
		}
	}()

	return doneCh
}

// fan-out - fan-in pattern.
func (a *App) createChan(ctx context.Context) InMetricChan {
	streams := make([]InMetricChan, len(a.parsers))
	for i, parFn := range a.parsers {
		p := parFn()
		select {
		case <-ctx.Done():
			break
		default:
		}
		streams[i] = a.result2Metric(ctx, p.Parse(ctx))
	}

	return a.muxChannels(ctx, streams...)
}

func (a *App) result2Metric(ctx context.Context, inCh <-chan collector.Result) InMetricChan {
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
					a.log.Error("error in data", zap.Error(r.Err))

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

func (a *App) muxChannels(ctx context.Context, streams ...InMetricChan) InMetricChan {
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
