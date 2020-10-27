package parsers

import (
	"context"

	"github.com/balabanovds/system-monitor/internal/collector"
)

type Parser interface {
	Parse(ctx context.Context) <-chan collector.Result
}
