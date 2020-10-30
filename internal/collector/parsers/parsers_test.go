package parsers_test

import (
	"context"
	"errors"
	"testing"

	"github.com/balabanovds/system-monitor/internal/collector/command"
	"github.com/balabanovds/system-monitor/internal/collector/parsers"
	"github.com/balabanovds/system-monitor/internal/models"
	"github.com/stretchr/testify/require"
)

func TestParsers(t *testing.T) {
	tests := []struct {
		name   string
		pType  models.ParserType
		expLen int
		err    error
	}{
		{
			name:   "load_avg",
			pType:  models.LoadAvg,
			expLen: 3,
		},
		{
			name:   "cpu",
			pType:  models.CPU,
			expLen: 3,
		},
		{
			name:  "not_found",
			pType: models.Undef,
			err:   command.ErrCommandNotFound,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			p, err := parsers.New(tt.pType)
			if tt.err != nil {
				require.Error(t, err)
				require.True(t, errors.As(err, &tt.err))

				return
			}

			require.NoError(t, err)

			result := make([]parsers.Result, 0)

			for res := range p.Parse(context.TODO()) {
				result = append(result, res)
			}

			require.Len(t, result, tt.expLen)
		})
	}
}
