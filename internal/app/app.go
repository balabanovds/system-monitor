package app

import (
	"context"
	"time"

	"github.com/balabanovds/system-monitor/internal/models"
)

type App interface {
	Run(ctx context.Context) <-chan struct{}
	RequestParsersInfo() []ParserInfo
	RequestStream(ctx context.Context, n, m int) InMetricChan
	GetMaxMeasurementsDuration() time.Duration
}

type InMetricChan <-chan models.Metric
