package app

import (
	"testing"

	"github.com/balabanovds/system-monitor/cmd/config"
	"github.com/balabanovds/system-monitor/internal/models"
	"github.com/balabanovds/system-monitor/internal/storage/inmem"
	"go.uber.org/zap"
)

func NewTestParsers() []models.ParserType {
	return []models.ParserType{models.LoadAvg, models.CPU}
}

func NewTestApp(t *testing.T, parsers []models.ParserType) *App {
	t.Helper()

	cfg := config.AppConfig{
		IntervalSeconds:     1,
		Parsers:             parsers,
		MaxMeasurementHours: 1,
	}

	logger := zap.NewNop()

	return New(cfg, inmem.New(logger), logger)
}
