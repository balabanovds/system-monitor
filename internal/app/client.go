package app

import (
	"context"
	"math"
	"time"

	"github.com/balabanovds/smonitor/internal/models"
)

// Request each n seconds provide report for last m seconds.
func (a *App) Request(ctx context.Context, n, m int) InMetricChan {
	outCh := make(chan models.Metric)

	if n < m {
		time.Sleep(time.Duration(m-n) * time.Second)
	}

	go func() {
		ticker := time.NewTicker(time.Duration(n) * time.Second)
		defer func() {
			ticker.Stop()
			close(outCh)
		}()

		for {
			select {
			case <-ctx.Done():
				return
			case t := <-ticker.C:
				for _, m := range report(a.storage.Get(t, time.Duration(m)*time.Second)) {
					select {
					case <-ctx.Done():
						return
					case outCh <- m:
					}
				}
			}
		}
	}()

	return outCh
}

func report(list []models.Metric) []models.Metric {
	mmap := make(map[models.MetricType][]models.Metric)

	for _, m := range list {
		mmap[m.Type] = append(mmap[m.Type], m)
	}

	result := make([]models.Metric, 0)

	for _, ms := range mmap {
		result = append(result, avg(ms))
	}
	return result
}

func avg(list []models.Metric) models.Metric {
	l := len(list)
	if l == 0 {
		return models.Metric{}
	}

	res := list[l-1]

	sum := .0
	for _, m := range list {
		sum += m.Value
	}
	res.Value = math.Round(sum*100) / 100

	return res
}
