package parsers

import (
	"context"
	"fmt"
	"os/exec"
	"strconv"
	"time"

	"github.com/balabanovds/smonitor/internal/models"

	"github.com/balabanovds/smonitor/internal/collector"
)

type LoadAvgParser struct {
	col *collector.Collector
}

func NewLoadAvgParser() Parser {
	return &LoadAvgParser{
		col: collector.New(
			exec.Command("uptime"),
		),
	}
}

func (p *LoadAvgParser) String() string {
	return "load_average"
}

func (p *LoadAvgParser) Parse(ctx context.Context) <-chan collector.Result {
	return p.col.Run(ctx, p.parseFn)
}

func (p *LoadAvgParser) parseFn(data []byte) ([]models.Metric, error) {
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
			Type:  models.LoadAverage1Min,
			Title: "load average 1 min",
			Value: ffields[0],
		},
		{
			Time:  now,
			Type:  models.LoadAverage5Min,
			Title: "load average 5 min",
			Value: ffields[1],
		},
		{
			Time:  now,
			Type:  models.LoadAverage15Min,
			Title: "load average 15 min",
			Value: ffields[2],
		},
	}, nil
}
