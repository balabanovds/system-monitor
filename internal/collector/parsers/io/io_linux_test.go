package io_test

import (
	"testing"

	"github.com/balabanovds/system-monitor/internal/collector/parsers/io"
	"github.com/balabanovds/system-monitor/internal/models"
	"github.com/stretchr/testify/require"
)

func TestParserFunc(t *testing.T) {
	data := `
{"sysstat": {
 	"hosts": [
 		{
 			"nodename": "BDT233",
 			"sysname": "Linux",
 			"release": "5.4.0-52-generic",
 			"machine": "x86_64",
 			"number-of-cpus": 8,
 			"date": "30.10.2020",
 			"statistics": [
 				{
 					"avg-cpu":  {"user": 10.67, "nice": 0.00, "system": 5.02, "iowait": 0.26, "steal": 0.00, "idle": 84.05},
 					"disk": [
 						{"disk_device": "loop0", "tps": 0.00, "kB_read/s": 0.00, "kB_wrtn/s": 0.00, "kB_read": 44, "kB_wrtn": 0},
 						{"disk_device": "loop7", "tps": 0.00, "kB_read/s": 0.00, "kB_wrtn/s": 0.00, "kB_read": 46, "kB_wrtn": 0},
 						{"disk_device": "nvme0n1", "tps": 213.21, "kB_read/s": 68.44, "kB_wrtn/s": 1832.81, "kB_read": 5958762, "kB_wrtn": 159566317}
 					]
 				}
 			]
 		}
 	]
 }}
`

	metrics, err := io.ParserFunc([]byte(data))
	require.NoError(t, err)

	ioTps := models.FilterByMetricType(metrics, models.IOtps)
	require.Len(t, ioTps, 1)
	require.Equal(t, 213.21, ioTps[0].Value)

	ioReadKbps := models.FilterByMetricType(metrics, models.IOReadKbps)
	require.Len(t, ioReadKbps, 1)
	require.Equal(t, 68.44, ioReadKbps[0].Value)

	ioWriteKbps := models.FilterByMetricType(metrics, models.IOWriteKbps)
	require.Len(t, ioWriteKbps, 1)
	require.Equal(t, 1832.81, ioWriteKbps[0].Value)

	ioCPUidle := models.FilterByMetricType(metrics, models.IOCPUidle)
	require.Len(t, ioCPUidle, 1)
	require.Equal(t, 84.05, ioCPUidle[0].Value)

	ioCPUuser := models.FilterByMetricType(metrics, models.IOCPUuser)
	require.Len(t, ioCPUuser, 1)
	require.Equal(t, 10.67, ioCPUuser[0].Value)

	ioCPUsystem := models.FilterByMetricType(metrics, models.IOCPUsystem)
	require.Len(t, ioCPUsystem, 1)
	require.Equal(t, 5.02, ioCPUsystem[0].Value)
}
