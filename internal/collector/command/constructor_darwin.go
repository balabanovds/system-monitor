package command

import (
	"github.com/balabanovds/system-monitor/internal/models"
)

func New() Command {
	return &commander{
		data: map[models.ParserType]string{
			models.LoadAvg: `uptime`,
			models.CPU:     `top -l 2 -n 0 | egrep '^CPU usage'`,
			models.IO:      `iostat`,
		},
	}
}
