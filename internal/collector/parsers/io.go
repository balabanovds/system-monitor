package parsers

import (
	"context"

	"github.com/balabanovds/smonitor/internal/models"

	"github.com/balabanovds/smonitor/internal/collector"
)

type IOParser struct {
	col *collector.Collector
}

func NewIOParser() Parser {
	return &IOParser{
		col: &collector.Collector{},
	}
}

func (p *IOParser) Parse(ctx context.Context) <-chan collector.Result {
	return p.col.Run(ctx, p.parseFn)
}

func (p *IOParser) parseFn(data []byte) ([]models.Metric, error) {
	// TODO complete
	return nil, nil
}
