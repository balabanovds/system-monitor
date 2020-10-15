package cmd_test

import (
	"bytes"
	"context"
	"os/exec"
	"testing"

	c "github.com/balabanovds/smonitor/pkg/cmd"

	"github.com/stretchr/testify/require"
)

func TestPipe(t *testing.T) {
	echoCmd := exec.Command("echo", "-n", "test")
	trCmd := exec.Command("tr", "'a-z'", "'A-Z'")
	pipe, err := c.New(echoCmd, trCmd)
	require.NoError(t, err)
	var out bytes.Buffer
	err = pipe.Run(context.TODO(), &out)
	require.NoError(t, err)
	require.Equal(t, "TEST", out.String())
}
