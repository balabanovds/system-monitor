package app_test

import (
	"context"
	"testing"
	"time"

	"github.com/balabanovds/system-monitor/internal/app"
	"github.com/balabanovds/system-monitor/internal/models"
	"github.com/stretchr/testify/require"
)

func TestApp_RequestParsersInfo(t *testing.T) {
	tests := []struct {
		name     string
		parsers  []models.ParserType
		length   int
		expected []app.ParserInfo
	}{
		{
			name:     "no parsers",
			length:   0,
			parsers:  []models.ParserType{},
			expected: []app.ParserInfo{},
		},
		{
			name:    "defined parsers",
			parsers: app.NewTestParsers(),
			length:  2,
			expected: []app.ParserInfo{
				{
					Type:        models.LoadAvg,
					Name:        "load_avg",
					MetricTypes: []models.MetricType{models.LoadAverage1Min, models.LoadAverage5Min, models.LoadAverage15Min},
				},
				{
					Type:        models.CPU,
					Name:        "cpu",
					MetricTypes: []models.MetricType{models.CPUUser, models.CPUSystem, models.CPUIdle},
				},
			},
		},
	}

	for _, tst := range tests {
		tst := tst
		t.Run(tst.name, func(t *testing.T) {
			a := app.NewTestApp(t, tst.parsers)
			infoList := a.RequestParsersInfo()
			require.Len(t, infoList, tst.length)
			require.Equal(t, tst.expected, infoList)
		})
	}
}

func TestApp_RequestStream(t *testing.T) {
	a := app.NewTestApp(t, app.NewTestParsers())

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	go func() {
		// run collection for 10 seconds
		<-a.Run(ctx)
	}()

	time.Sleep(2500 * time.Millisecond)
	metrics := make([]models.Metric, 0)
	for m := range a.RequestStream(ctx, 2, 3) {
		metrics = append(metrics, m)
	}

	// we got 3 ticks, each 6 metrics
	require.Len(t, metrics, 18)
}
