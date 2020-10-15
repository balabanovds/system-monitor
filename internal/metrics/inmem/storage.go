package inmem

import (
	"sync"
	"time"

	"github.com/balabanovds/smonitor/internal/models"

	"github.com/balabanovds/smonitor/internal/metrics"
)

type storage struct {
	mu   sync.RWMutex
	data map[time.Time][]models.Metric
}

func New() metrics.Storage {
	return &storage{
		data: make(map[time.Time][]models.Metric),
	}
}

func (s *storage) Get(end time.Time, duration time.Duration) []models.Metric {
	var result []models.Metric
	s.mu.RLock()
	defer s.mu.RUnlock()

	start := end.Add(-duration)

	for k, v := range s.data {
		if k.UTC().UnixNano() > start.UTC().UnixNano() &&
			k.UTC().UnixNano() < end.UTC().UnixNano() {
			result = append(result, v...)
		}
	}

	return result
}

func (s *storage) Save(m models.Metric) {
	s.mu.Lock()
	defer s.mu.Unlock()
	d, ok := s.data[m.Time]
	if !ok {
		s.data[m.Time] = []models.Metric{m}
		return
	}

	s.data[m.Time] = append(d, m)
}

func (s *storage) Delete(till time.Time) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for k := range s.data {
		if k.UTC().UnixNano() < till.UTC().UnixNano() {
			delete(s.data, k)
		}
	}
}
