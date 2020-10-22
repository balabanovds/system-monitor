package collector

import (
	"bytes"
	"context"
	"os/exec"

	"github.com/balabanovds/smonitor/internal/models"
)

type Collector struct {
	cmd *exec.Cmd
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

func New(command string) *Collector {
	return &Collector{
		cmd: exec.Command("/bin/sh", "-c", command),
	}
}

// pipeline pattern.
func (c *Collector) Run(ctx context.Context, parseFn ParseFn) <-chan Result {
	return c.parse(ctx, c.execCmd(ctx, c.cmd), parseFn)
}

func (c *Collector) execCmd(ctx context.Context, cmd *exec.Cmd) <-chan ExecResult {
	stream := make(chan ExecResult)
	go func() {
		defer close(stream)

		var out bytes.Buffer
		cmd.Stdout = &out

		if err := cmd.Run(); err != nil {
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
			metrics, err := parseFn(i.Data)
			if err != nil {
				select {
				case <-ctx.Done():
					return
				case stream <- Result{Err: i.Err}:
				}
				continue
			}

			for _, m := range metrics {
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
