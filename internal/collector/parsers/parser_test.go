package parser_test

import (
	"context"
	"errors"
	"testing"

	"github.com/balabanovds/system-monitor/internal/collector/command"
	parser "github.com/balabanovds/system-monitor/internal/collector/parsers"
	"github.com/balabanovds/system-monitor/internal/models"
	"github.com/stretchr/testify/require"
)

func TestParsers(t *testing.T) {
	tests := []struct {
		name        string
		parserType  models.ParserType
		expectedLen int
		err         error
	}{
		{
			name:        "load_avg",
			parserType:  models.LoadAvg,
			expectedLen: 3,
		},
		{
			name:        "cpu",
			parserType:  models.CPU,
			expectedLen: 3,
		},
		{
			name:       "not_found",
			parserType: models.Undef,
			err:        command.ErrCommandNotImplemented,
		},
		{
			name:       "io",
			parserType: models.IO,
			err:        parser.ErrFailedRunCommand,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			p, err := parser.New(tt.parserType)
			if err != nil {
				require.True(t, errors.As(err, &tt.err))

				return
			}

			require.NoError(t, err)

			result := make([]parser.Result, 0)

			for res := range p.Parse(context.TODO()) {
				if res.Err != nil {
					require.EqualError(t, res.Err, tt.err.Error())

					return
				}
				result = append(result, res)
			}

			require.Len(t, result, tt.expectedLen)
		})
	}
}
