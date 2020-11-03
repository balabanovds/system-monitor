package parser

import (
	"context"
	"errors"
	"fmt"
	"os/exec"

	"github.com/balabanovds/system-monitor/internal/collector/command"
	pcpu "github.com/balabanovds/system-monitor/internal/collector/parsers/cpu"
	pfs "github.com/balabanovds/system-monitor/internal/collector/parsers/fs"
	pio "github.com/balabanovds/system-monitor/internal/collector/parsers/io"
	pavg "github.com/balabanovds/system-monitor/internal/collector/parsers/load_avg"
	pnet "github.com/balabanovds/system-monitor/internal/collector/parsers/net"
	"github.com/balabanovds/system-monitor/internal/models"
)

var (
	ErrParserNotFound   = errors.New("parser not found")
	ErrFailedRunCommand = errors.New("failed to start command")
)

type Parser interface {
	Parse(ctx context.Context) <-chan Result
	Type() models.ParserType
	Error() error
}

type Func func([]byte) ([]models.Metric, error)

type parser struct {
	pType models.ParserType
	pFunc Func
	cmd   *exec.Cmd
}

func New(t models.ParserType) (Parser, error) {
	cmd, err := command.New().Get(t)
	if err != nil {
		return nil, fmt.Errorf("failed init %s parser: %w", t, err)
	}
	pFunc := getParserFunc(t)
	if pFunc == nil {
		return nil, ErrParserNotFound
	}

	return &parser{
		pType: t,
		pFunc: pFunc,
		cmd:   cmd,
	}, nil
}

// pipeline pattern.
func (p *parser) Parse(ctx context.Context) <-chan Result {
	return parseExecResult(ctx, execCmd(ctx, p.cmd), p.pFunc)
}

func (p *parser) Type() models.ParserType {
	return p.pType
}

func (p *parser) Error() error {
	return getError(p.pType)
}

func getParserFunc(t models.ParserType) Func {
	switch t {
	case models.LoadAvg:
		return pavg.ParserFunc
	case models.CPU:
		return pcpu.ParserFunc
	case models.IO:
		return pio.ParserFunc
	case models.FS:
		return pfs.ParserFunc
	case models.Net:
		return pnet.ParserFunc
	case models.Undef:
	}

	return nil
}

func getError(t models.ParserType) error {
	switch t {
	case models.LoadAvg:
		return errors.New(`load_avg failed to run: pls consider if 'uptime' is installed on OS`)
	case models.CPU:
		return errors.New(`cpu failed to run: pls consider if 'top' is installed on OS`)
	case models.IO:
		return errors.New(`io failed to run: pls consider if 'iostat' is installed on OS`)
	case models.FS:
		// TODO fill
		return errors.New(`TBD`)
	case models.Net:
		// TODO fill
		return errors.New(`TBD`)
	case models.Undef:
	}

	return nil
}
