package parsers

import (
	"context"

	"github.com/balabanovds/system-monitor/internal/collector"
	"github.com/balabanovds/system-monitor/internal/models"
)

type NETParser struct {
	col *collector.Collector
}

func NewNETParser() Parser {
	return &NETParser{
		col: &collector.Collector{},
	}
}

func (p *NETParser) Parse(ctx context.Context) <-chan collector.Result {
	return p.col.Run(ctx, p.parseFn)
}

func (p *NETParser) parseFn(data []byte) ([]models.Metric, error) {
	// TODO complete
	return nil, nil
}
