package command

import (
	"github.com/balabanovds/system-monitor/internal/models"
)

func New() Command {
	return &commander{
		data: map[models.ParserType]string{
			models.LoadAvg: "uptime",
			models.CPU:     `top -i -b -n1 | egrep '^%Cpu'`,
		},
	}
}
