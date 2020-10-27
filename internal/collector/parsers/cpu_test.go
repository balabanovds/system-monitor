package parsers_test

import (
	"context"
	"testing"

	"github.com/balabanovds/system-monitor/internal/collector/parsers"
	"github.com/stretchr/testify/require"
)

func TestCPU(t *testing.T) {
	parser := parsers.NewCPUParser()
	dataCh := parser.Parse(context.TODO())

	var data []float64 //nolint:prealloc
	for d := range dataCh {
		require.NoError(t, d.Err)
		data = append(data, d.Data.Value)
	}
	require.Len(t, data, 3)
}
