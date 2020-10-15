package models

import (
	"time"
)

type MetricType uint8

const (
	Undefined MetricType = iota
	LoadAverage1Min
	LoadAverage5Min
	LoadAverage15Min
	CPUUser
	CPUSystem
	CPUIdle
	IOtps
	IOReadKbps
	IOWriteKbps
	IOCPUuser
	IOCPUsystem
	IOCPUidle
	FSMBFree
	FSInodeFree
)

type Metric struct {
	Time  time.Time
	Type  MetricType
	Title string
	Value float64
}
