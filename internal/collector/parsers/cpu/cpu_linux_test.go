package cpu_test

import (
	"testing"

	"github.com/balabanovds/system-monitor/internal/collector/parsers/cpu"
	"github.com/stretchr/testify/require"
)

func TestParserFunc(t *testing.T) {
	data := `%Cpu(s): 10,7 us,  3,9 sy,  0,0 ni, 84,0 id,  0,3 wa,  0,0 hi,  1,1 si,  0,0 st`

	metrics, err := cpu.ParserFunc([]byte(data))
	require.NoError(t, err)
	require.Equal(t, 10.7, metrics[0].Value)
	require.Equal(t, 3.9, metrics[1].Value)
	require.Equal(t, 84.0, metrics[2].Value)
}
