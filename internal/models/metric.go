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
	case Undef:
	}

	return []MetricType{Undefined}
}

func (m *Metric) IsWithinTime(start, stop time.Time) bool {
	if start.UTC().UnixNano() > stop.UTC().UnixNano() {
		return false
	}

	return m.Time.UTC().UnixNano() >= start.UTC().UnixNano() &&
		m.Time.UTC().UnixNano() <= stop.UTC().UnixNano()
}

func (m *Metric) IsLessTime(t time.Time) bool {
	return m.Time.UTC().UnixNano() < t.UTC().UnixNano()
}

func (m *Metric) IsLess(other Metric) bool {
	if m.Time.UTC().UnixNano() == other.Time.UTC().UnixNano() {
		return m.Type < other.Type
	}

	return m.Time.UTC().UnixNano() < other.Time.UTC().UnixNano()
}
