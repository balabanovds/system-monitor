package parsers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/balabanovds/smonitor/internal/models"

	"github.com/balabanovds/smonitor/internal/collector"
)

type CPUParser struct {
	col *collector.Collector
}

func (p *CPUParser) Parse(ctx context.Context) <-chan collector.Result {
	return p.col.Run(ctx, p.parseFn)
}

func (p *CPUParser) parseFn(data []byte) ([]models.Metric, error) {
	now := time.Now()
	fields := p.parse(string(data))

	ffields := make([]float64, 0)
	for _, f := range fields {
		fl, err := strconv.ParseFloat(f, 64)
		if err != nil {
			return nil, err
		}
		ffields = append(ffields, fl)
	}

	if len(ffields) < 3 {
		return nil, fmt.Errorf("parsing failed")
	}

	return []models.Metric{
		{
			Time:  now,
			Type:  models.CPUUser,
			Title: "%user",
			Value: ffields[0],
		},
		{
			Time:  now,
			Type:  models.CPUSystem,
			Title: "%system",
			Value: ffields[1],
		},
		{
			Time:  now,
			Type:  models.CPUIdle,
			Title: "%idle",
			Value: ffields[2],
		},
	}, nil
}
