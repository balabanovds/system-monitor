package loadavg

import (
	"strings"
)

func parse(str string) []string {
	spl := strings.Split(str, "load averages:")
	fields := strings.Fields(spl[1])

	return []string{fields[0], fields[1], fields[2]}
}
