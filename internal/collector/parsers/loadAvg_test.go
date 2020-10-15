package parsers_test

import (
	"context"
	"testing"

	"github.com/balabanovds/smonitor/internal/collector/parsers"

	"github.com/stretchr/testify/require"
)

func TestLoadAvg(t *testing.T) {
	parser := parsers.NewLoadAvgParser()
	dataCh := parser.Parse(context.TODO())

	var data []float64 //nolint:prealloc
	for d := range dataCh {
		require.NoError(t, d.Err)
		data = append(data, d.Data.Value)
	}
	require.Len(t, data, 3)
}
