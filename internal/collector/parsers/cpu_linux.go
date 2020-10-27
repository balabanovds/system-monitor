package parsers

import (
	"bytes"
	"strings"

	"github.com/balabanovds/system-monitor/internal/collector"
)

func NewCPUParser() Parser {
	// top -i -b -n1 | egrep '^%Cpu'
	return &CPUParser{
		col: collector.New(`top -i -b -n1 | egrep '^%Cpu'`),
	}
}

func (p *CPUParser) parse(data []byte) ([]string, error) {
	// %Cpu(s):  0.2 us,  0.1 sy,  0.0 ni, 99.6 id,  0.0 wa,  0.0 hi,  0.0 si,  0.0 st

	line, err := readLine(bytes.NewReader(data), 1)
	if err != nil {
		return nil, err
	}
	fields := strings.Fields(line)

	return []string{
		strings.ReplaceAll(fields[1], ",", "."),
		strings.ReplaceAll(fields[3], ",", "."),
		strings.ReplaceAll(fields[7], ",", "."),
	}, nil
}
