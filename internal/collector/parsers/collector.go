package parser

import (
	"bytes"
	"context"
	"os/exec"

	"github.com/balabanovds/system-monitor/internal/models"
)

type Result struct {
	Data models.Metric
	Err  error
}

type ExecResult struct {
	Data []byte
	Err  error
}

func execCmd(ctx context.Context, cmd *exec.Cmd) <-chan ExecResult {
	stream := make(chan ExecResult)
	go func() {
		defer close(stream)

		var out bytes.Buffer
		cmd.Stdout = &out

		if err := cmd.Run(); err != nil {
			stream <- ExecResult{Err: ErrFailedRunCommand}

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

func parseExecResult(ctx context.Context, inStream <-chan ExecResult, parseFn Func) <-chan Result {
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
