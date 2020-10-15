package parsers

import (
	"os/exec"
	"strings"

	"github.com/balabanovds/smonitor/internal/collector"
)

func NewCPUParser() Parser {
	return &CPUParser{
		col: collector.New(
			exec.Command("top", "-l", "2", "-n", "0"),
			exec.Command("egrep", "'^CPU usage'"),
			exec.Command("tail", "-n1"),
		),
	}
}

func (p *CPUParser) parse(str string) []string {
	// CPU usage: 9.0% user, 11.20% sys, 79.80% idle
	fields := strings.Fields(str)
	return []string{
		strings.Trim(fields[2], "%"),
		strings.Trim(fields[4], "%"),
		strings.Trim(fields[6], "%"),
	}
}
