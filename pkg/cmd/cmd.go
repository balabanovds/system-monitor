package cmd

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

type CmdPipe struct {
	commands []*exec.Cmd
}

func New(commands ...*exec.Cmd) (*CmdPipe, error) {
	if len(commands) == 0 {
		return nil, fmt.Errorf("no commands")
	}
	for i := range commands {
		commands[i].Stderr = os.Stderr
		if i > 0 {
			var err error
			if commands[i].Stdin, err = commands[i-1].StdoutPipe(); err != nil {
				return nil, fmt.Errorf("pipe commands %s -> %s failed: %w",
					commands[i-1], commands[i], err)
			}
		}
	}

	return &CmdPipe{
		commands: commands,
	}, nil
}

func (p *CmdPipe) Run(ctx context.Context, out io.Writer) error {
	defer func() {
		for _, err := range p.stop(ctx) {
			log.Println(err)
		}
	}()
	p.commands[len(p.commands)-1].Stdout = out
	for _, c := range p.commands {
		select {
		case <-ctx.Done():
			return nil
		default:
			if err := c.Start(); err != nil {
				return fmt.Errorf("pipe: command %s failed to start: %w", c, err)
			}
		}

	}
	return nil
}

func (p *CmdPipe) stop(ctx context.Context) []error {
	var errs []error
	for _, c := range p.commands {
		select {
		case <-ctx.Done():
			return errs
		default:
			if err := c.Wait(); err != nil {
				errs = append(errs, err)
			}
		}

	}
	return errs
}

func RunPipe(ctx context.Context, cmds ...*exec.Cmd) OutCh {
	var stages []stage
	for _, c := range cmds {
		stages = append(stages, newStage(c))
	}

	inCh := make(BiCh)
	out := execPipeline(ctx, inCh, stages...)
	inCh <- Result{}

	return out
}

func newStage(c *exec.Cmd) stage {
	return func(inCh InCh) OutCh {
		outCh := make(BiCh)
		go func() {
			defer close(outCh)
			for v := range inCh {
				if v.Err != nil {
					continue
				}
				var out bytes.Buffer
				c.Stdout = &out
				c.Stdin = bytes.NewBuffer(v.Data)
				if err := c.Start(); err != nil {
					outCh <- Result{Err: err}
					continue
				}
				if err := c.Wait(); err != nil {
					outCh <- Result{Err: err}
					continue
				}
				outCh <- Result{Data: out.Bytes(), Err: nil}
			}
		}()
		return outCh
	}
}
