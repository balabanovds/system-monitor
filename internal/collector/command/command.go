package command

import (
	"errors"
	"os/exec"

	"github.com/balabanovds/system-monitor/internal/models"
)

type Command interface {
	Get(pType models.ParserType) (*exec.Cmd, error)
}

var ErrCommandNotFound = errors.New("command not found")

type commander struct {
	data map[models.ParserType]string
}

func (c *commander) Get(pType models.ParserType) (*exec.Cmd, error) {
	cmdStr, ok := c.data[pType]
	if !ok {
		return nil, ErrCommandNotFound
	}

	return exec.Command("/bin/sh", "-c", cmdStr), nil
}
