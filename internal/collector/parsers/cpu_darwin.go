package parsers

import (
	"bytes"
	"strings"

	"github.com/balabanovds/system-monitor/internal/collector"
)

func NewCPUParser() Parser {
	return &CPUParser{
		col: collector.New(`top -l 2 -n 0 | egrep '^CPU usage'`),
	}
}

func (p *CPUParser) parse(data []byte) ([]string, error) {
	// expected two lines, should take second
	// CPU usage: 9.0% user, 11.20% sys, 79.80% idle
	// CPU usage: 9.0% user, 11.20% sys, 79.80% idle
	line, err := readLine(bytes.NewReader(data), 2)
	if err != nil {
		return nil, err
	}

	fields := strings.Fields(line)
	return []string{
		strings.Trim(fields[2], "%"),
		strings.Trim(fields[4], "%"),
		strings.Trim(fields[6], "%"),
	}, nil
}
