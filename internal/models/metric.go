package models

import (
	"time"
)

type Metric struct {
	Time  time.Time
	Type  MetricType
	Title string
	Value float64
}

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

func GetMetricTypes(pType ParserType) []MetricType {
	switch pType {
	case LoadAvg:
		return []MetricType{LoadAverage1Min, LoadAverage5Min, LoadAverage15Min}
	case CPU:
		return []MetricType{CPUUser, CPUSystem, CPUIdle}
	case IO:
		return []MetricType{IOtps, IOReadKbps, IOWriteKbps, IOCPUuser, IOCPUsystem, IOCPUidle}
	case FS:
		return []MetricType{FSMBFree, FSInodeFree}
	case Net:
		// TODO complete Net parser here too
	}
	return []MetricType{Undefined}
}
