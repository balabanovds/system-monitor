package collector

import (
	"bytes"
	"context"
	"os/exec"

	"github.com/balabanovds/smonitor/internal/models"

	"github.com/balabanovds/smonitor/pkg/cmd"
)

type Collector struct {
	cmds []*exec.Cmd
}

type Result struct {
	Data models.Metric
	Err  error
}

type ExecResult struct {
	Data []byte
	Err  error
}

type ParseFn func(data []byte) ([]models.Metric, error)

func New(cmds ...*exec.Cmd) *Collector {
	return &Collector{cmds}
}

// pipeline pattern.
func (c *Collector) Run(ctx context.Context, parseFn ParseFn) <-chan Result {
	return c.parse(ctx, c.execCmd(ctx, c.cmds...), parseFn)
}

func (c *Collector) execCmd(ctx context.Context, cmds ...*exec.Cmd) <-chan ExecResult {
	stream := make(chan ExecResult)
	go func() {
		defer close(stream)

		p, err := cmd.New(cmds...)
		if err != nil {
			stream <- ExecResult{Err: err}
			return
		}

		var out bytes.Buffer
		if err != p.Run(ctx, &out) {
			stream <- ExecResult{Err: err}
			return
		}

		select {
		case <-ctx.Done():
			return
		case stream <- ExecResult{Data: out.Bytes()}:
		}
	}()

	return stream
}

func (c *Collector) parse(ctx context.Context, inStream <-chan ExecResult, parseFn ParseFn) <-chan Result {
	stream := make(chan Result)
	go func() {
		defer close(stream)
		for i := range inStream {
			if i.Err != nil {
				select {
				case <-ctx.Done():
					return
				case stream <- Result{Err: i.Err}:
				}
				continue
			}
			mtrcs, err := parseFn(i.Data)
			if err != nil {
				select {
				case <-ctx.Done():
					return
				case stream <- Result{Err: i.Err}:
				}
				continue
			}

			for _, m := range mtrcs {
				select {
				case <-ctx.Done():
					return
				case stream <- Result{Data: m}:
				}
			}
		}
	}()

	return stream
}
