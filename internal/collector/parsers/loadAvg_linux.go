package parsers

import "strings"

func (p *LoadAvgParser) parse(str string) []string {
	spl := strings.Split(str, "load average:")
	fields := strings.Fields(spl[1])

	return []string{
		strings.ReplaceAll(strings.Trim(fields[0], ","), ",", "."),
		strings.ReplaceAll(strings.Trim(fields[2], ","), ",", "."),
		strings.ReplaceAll(strings.Trim(fields[1], ","), ",", "."),
	}
}
