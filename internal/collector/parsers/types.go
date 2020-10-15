package parsers

import (
	"context"

	"github.com/balabanovds/smonitor/internal/collector"
)

type Parser interface {
	Parse(ctx context.Context) <-chan collector.Result
}
