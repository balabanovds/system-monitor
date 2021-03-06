package app

import (
	"github.com/balabanovds/system-monitor/internal/models"
)

type ParserInfo struct {
	Type        models.ParserType
	Name        string
	MetricTypes []models.MetricType
}

func newParserInfo(ptype models.ParserType) ParserInfo {
	return ParserInfo{
		Type:        ptype,
		Name:        ptype.String(),
		MetricTypes: models.GetMetricTypes(ptype),
	}
}

func (a *Monitor) RequestParsersInfo() []ParserInfo {
	list := make([]ParserInfo, len(a.parserTypes))

	for i, t := range a.parserTypes {
		list[i] = newParserInfo(t)
	}

	return list
}
