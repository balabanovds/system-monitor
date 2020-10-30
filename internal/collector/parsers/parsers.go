package parsers

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

var ErrorParserNotFound = errors.New("parser not found")

type Parser interface {
	Parse(ctx context.Context) <-chan Result
	Type() models.ParserType
}

type ParserFunc func([]byte) ([]models.Metric, error)

type parser struct {
	pType models.ParserType
	pFunc ParserFunc
	cmd   *exec.Cmd
}

func New(t models.ParserType) (Parser, error) {
	cmd, err := command.New().Get(t)
	if err != nil {
		return nil, fmt.Errorf("failed init %s parser: %w", t, err)
	}
	pFunc := getParserFunc(t)
	if pFunc == nil {
		return nil, ErrorParserNotFound
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

func getParserFunc(t models.ParserType) ParserFunc {
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
