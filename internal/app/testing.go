package app

import (
	"github.com/balabanovds/system-monitor/cmd/config"
	"github.com/balabanovds/system-monitor/internal/models"
	"github.com/balabanovds/system-monitor/internal/storage/inmem"
	"go.uber.org/zap"
	"testing"
)

func NewTestApp(t *testing.T, parser ...models.ParserType) *App {
	t.Helper()

	parsers := []models.ParserType{models.LoadAvg, models.CPU}
	newParsers := make([]models.ParserType, 0)
OUTER:
	for _, np := range parser {
		for _, op := range parsers {
			if op == np {
				continue OUTER
			}
			newParsers = append(newParsers, np)
		}
	}

	parsers = append(parsers, newParsers...)

	cfg := config.AppConfig{
		IntervalSeconds: 1,
		Parsers:         parsers,
	}

	logger := zap.NewNop()

	return New(cfg, inmem.New(logger), logger)
}
