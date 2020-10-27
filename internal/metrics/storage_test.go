package metrics_test

import (
	"testing"
	"time"

	"github.com/balabanovds/system-monitor/internal/metrics/inmem"
	"github.com/balabanovds/system-monitor/internal/models"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

var logger = zap.NewNop()

func TestStorage_Get(t *testing.T) {
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

	mmap := st.Get(now.Add(time.Second), time.Minute)
	require.Len(t, mmap, 2)
}

func TestSTorage_Delete(t *testing.T) {
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
		Time: now.Add(2 * time.Second),
		Type: models.LoadAverage15Min,
	}

	st.Save(m1)
	st.Save(m2)
	st.Save(m3)

	st.Delete(now.Add(time.Second))

	mmap := st.Get(now.Add(3*time.Second), time.Minute)
	require.Len(t, mmap, 1)
}
