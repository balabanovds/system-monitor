package metrics

import (
	"time"

	"github.com/balabanovds/system-monitor/internal/models"
)

type Storage interface {
	// Get Metrics by MetricType within period
	Get(time time.Time, period time.Duration) []models.Metric
	Save(m models.Metric)
	// Delete all Metrics till time
	Delete(till time.Time)
}
