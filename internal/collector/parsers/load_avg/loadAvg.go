package loadavg

import (
	"fmt"
	"strconv"
	"time"

	"github.com/balabanovds/system-monitor/internal/models"
)

func ParserFunc(data []byte) ([]models.Metric, error) {
	now := time.Now()
	fields := parse(string(data))

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
