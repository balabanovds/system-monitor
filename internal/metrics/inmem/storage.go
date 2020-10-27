package inmem

import (
	"sync"
	"time"

	"github.com/balabanovds/system-monitor/internal/metrics"
	"github.com/balabanovds/system-monitor/internal/models"
	"go.uber.org/zap"
)

type storage struct {
	mu   sync.RWMutex
	data []models.Metric
	log  *zap.Logger
}

func New(logger *zap.Logger) metrics.Storage {
	return &storage{
		data: make([]models.Metric, 0),
		log:  logger,
	}
}

func (s *storage) Get(end time.Time, duration time.Duration) []models.Metric {
	s.mu.RLock()
	defer s.mu.RUnlock()

	start := end.Add(-duration)
	var idxStart int
	idxEnd := len(s.data)

	for i := idxEnd - 1; i >= 0; i-- {
		if !s.data[i].IsLessTime(end) {
			idxEnd = i
		}
		idxStart = i
		if s.data[i].IsLessTime(start) {
			break
		}
	}

	return s.data[idxStart:idxEnd]
}

func (s *storage) Save(m models.Metric) {
	s.mu.Lock()
	defer s.mu.Unlock()
	result := make([]models.Metric, 0)

	var idx int

	for i := len(s.data) - 1; i >= 0; i-- {
		idx = i + 1
		if !m.IsLess(s.data[i]) {
			break
		}
	}

	result = append(result, s.data[:idx]...)
	result = append(result, m)
	result = append(result, s.data[idx:]...)

	s.data = result
}

func (s *storage) Delete(till time.Time) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var idx int

	for i := range s.data {
		if s.data[i].IsLessTime(till) {
			idx = i
		}
	}

	s.data = s.data[idx:]
}
