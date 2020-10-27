package metrics_test

import (
	"testing"
	"time"

	"github.com/balabanovds/system-monitor/internal/metrics"
	"github.com/balabanovds/system-monitor/internal/metrics/inmem"
	"github.com/balabanovds/system-monitor/internal/models"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

var logger = zap.NewNop()

func TestStorage_Save(t *testing.T) {
	st, now := initStorage(t)

	m := models.Metric{
		Time: now.Add(2 * time.Second),
		Type: models.CPUIdle,
	}

	st.Save(m)

	list := st.Get(now.Add(time.Hour), 2*time.Hour)

	require.Len(t, list, 4)
	require.Equal(t, list[2], m)
}

func TestStorage_Get(t *testing.T) {
	st, now := initStorage(t)

	list := st.Get(now.Add(time.Second), time.Minute)
	require.Len(t, list, 2)
}

func TestStorage_Delete(t *testing.T) {
	st, now := initStorage(t)

	st.Delete(now.Add(time.Second))

	mmap := st.Get(now.Add(3*time.Second), time.Minute)
	require.Len(t, mmap, 1)
}

func initStorage(t *testing.T) (metrics.Storage, time.Time) {
	t.Helper()
	st := inmem.New(logger)

	now := time.Now()
	m1 := models.Metric{
		Time: now,
		Type: models.LoadAverage1Min,
	}

	m2 := models.Metric{
		Time: now,
		Type: models.LoadAverage15Min,
	}

	m3 := models.Metric{
		Time: now.Add(3 * time.Second),
		Type: models.LoadAverage15Min,
	}

	st.Save(m1)
	st.Save(m2)
	st.Save(m3)

	return st, now
}
