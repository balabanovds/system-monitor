package cpu_test

import (
	"testing"

	"github.com/balabanovds/system-monitor/internal/collector/parsers/cpu"
	"github.com/stretchr/testify/require"
)

func TestParserFunc(t *testing.T) {
	data := `CPU usage: 9.0% user, 11.20% sys, 79.80% idle
CPU usage: 9.0% user, 11.20% sys, 79.80% idle`

	metrics, err := cpu.ParserFunc([]byte(data))
	require.NoError(t, err)
	require.Equal(t, 9.0, metrics[0].Value)
	require.Equal(t, 11.2, metrics[1].Value)
	require.Equal(t, 79.8, metrics[2].Value)
}
