package parsers

import (
	"os/exec"
	"strings"

	"github.com/balabanovds/smonitor/internal/collector"
)

type CPUParser struct {
	col *collector.Collector
}

func NewCPUParser() Runnable {
	// top -i -b -n1 | egrep '^%Cpu'
	return &CPUParser{
		col: collector.New(
			exec.Command("top", "-i", "-b", "-n1"),
			exec.Command("egrep", "'^%Cpu'"),
		),
	}
}

func (p *CPUParser) parse(str string) []string {
	// %Cpu(s):  0.2 us,  0.1 sy,  0.0 ni, 99.6 id,  0.0 wa,  0.0 hi,  0.0 si,  0.0 st
	fields := strings.Fields(str)
	return []string{
		fields[1],
		fields[3],
		fields[7],
	}
}
