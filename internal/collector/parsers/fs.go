package parsers

import (
	"context"

	"github.com/balabanovds/smonitor/internal/models"

	"github.com/balabanovds/smonitor/internal/collector"
)

type FSParser struct {
	col *collector.Collector
}

func NewFSParser() Parser {
	return &FSParser{
		col: &collector.Collector{},
	}
}

func (p *FSParser) Parse(ctx context.Context) <-chan collector.Result {
	return p.col.Run(ctx, p.parseFn)
}

func (p *FSParser) parseFn(data []byte) ([]models.Metric, error) {
	// TODO complete
	return nil, nil
}
