//go:build !windows
// +build !windows

package term

import (
	"os"
	"os/exec"
)

func (c *Command) Make() *exec.Cmd {
	cmd := exec.Command(c.Bin, c.Args...)
	if !c.Hide {
		if c.Stdout {
			cmd.Stdout = os.Stdout
		}
		if c.Stderr {
			cmd.Stderr = os.Stderr
		}
		if c.Stdin {
			cmd.Stdin = os.Stdin
		}
	}
	return cmd
}
