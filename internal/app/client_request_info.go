package app

import "github.com/balabanovds/smonitor/internal/models"

type ParserInfo struct {
	Type        models.ParserType
	MetricTypes []models.MetricType
}

func newParserInfo(ptype models.ParserType) ParserInfo {
	return ParserInfo{
		Type:        ptype,
		MetricTypes: models.GetMetricTypes(ptype),
	}
}

func (a *App) RequestParsersInfo() []ParserInfo {
	var list []ParserInfo

	for _, t := range a.parserTypes {
		list = append(list, newParserInfo(t))
	}

	return list
}
