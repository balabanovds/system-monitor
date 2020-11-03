package app

import (
	"context"
	"sync"
	"time"

	"github.com/balabanovds/system-monitor/cmd/config"
	parser "github.com/balabanovds/system-monitor/internal/collector/parsers"
	"github.com/balabanovds/system-monitor/internal/models"
	"github.com/balabanovds/system-monitor/internal/storage"
	"go.uber.org/zap"
)

type Monitor struct {
	sync.Mutex
	storage                storage.Storage
	parserTypes            []models.ParserType
	interval               time.Duration
	maxMeasurementDuration time.Duration
	log                    *zap.Logger
}

func New(cfg config.AppConfig, storage storage.Storage, logger *zap.Logger) App {
	return &Monitor{
		storage:                storage,
		parserTypes:            cfg.Parsers,
		interval:               time.Duration(cfg.IntervalSeconds) * time.Second,
		maxMeasurementDuration: time.Duration(cfg.MaxMeasurementHours) * time.Hour,
		log:                    logger,
	}
}

func (a *Monitor) Run(ctx context.Context) <-chan struct{} {
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
				a.storage.Delete(t.Add(-a.maxMeasurementDuration))
				select {
				case <-ctx.Done():
					return
				default:
				}

				metricCh, err := a.createChan(ctx, a.parserTypes)
				if err != nil {
					a.log.Error("create channel failed", zap.Error(err))

					return
				}

				for m := range metricCh {
					a.storage.Save(m)
				}
			}
		}
	}()

	return doneCh
}

func (a *Monitor) GetMaxMeasurementsDuration() time.Duration {
	return a.maxMeasurementDuration
}

// fan-out - fan-in pattern.
func (a *Monitor) createChan(ctx context.Context, parserTypes []models.ParserType) (InMetricChan, error) {
	parserSlice := make([]parser.Parser, 0)

	for _, p := range parserTypes {
		pr, err := parser.New(p)
		if err != nil {
			return nil, err
		}
		parserSlice = append(parserSlice, pr)
	}

	streams := make([]InMetricChan, len(parserSlice))

	for i, p := range parserSlice {
		select {
		case <-ctx.Done():
			break
		default:
		}
		streams[i] = a.result2Metric(ctx, p.Parse(ctx))
	}

	return a.muxChannels(ctx, streams...), nil
}

func (a *Monitor) result2Metric(ctx context.Context, inCh <-chan parser.Result) InMetricChan {
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

func (a *Monitor) muxChannels(ctx context.Context, streams ...InMetricChan) InMetricChan {
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
