package cpu

import (
	"fmt"
	"strconv"
	"time"

	"github.com/balabanovds/system-monitor/internal/models"
)

func ParserFunc(data []byte) ([]models.Metric, error) {
	now := time.Now()
	fields, err := parse(data)
	if err != nil {
		return nil, err
	}

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
