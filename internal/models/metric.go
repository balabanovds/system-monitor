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

func (m MetricType) String() string {
	switch m {
	case Undefined:
		return "undefined"
	case LoadAverage1Min:
		return "load_avg_1min"
	case LoadAverage5Min:
		return "load_avg_5min"
	case LoadAverage15Min:
		return "load_avg_15min"
	case CPUUser:
		return "cpu_%user"
	case CPUSystem:
		return "cpu_%system"
	case CPUIdle:
		return "cpu_%idle"
	case IOtps:
		return "io_tps"
	case IOReadKbps:
		return "io_read_kbp/s"
	case IOWriteKbps:
		return "io_write_kbp/s"
	case IOCPUuser:
		return "io_cpu_%user"
	case IOCPUsystem:
		return "io_cpu_%system"
	case IOCPUidle:
		return "io_cpu_%idle"
	case FSMBFree:
		return "fs_%mb"
	case FSInodeFree:
		return "fs_%inode"
	}

	return ""
}

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

func FilterByMetricType(metrics []Metric, metricType MetricType) []Metric {
	res := make([]Metric, 0)

	for _, m := range metrics {
		if m.Type == metricType {
			res = append(res, m)
		}
	}

	return res
}
