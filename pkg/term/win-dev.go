//go:build windows
// +build windows

package term

import (
	"os"
	"os/exec"
	"syscall"
)

func (c *Command) Make() *exec.Cmd {
	cmd := exec.Command(c.Bin, c.Args...)
	// TODO:Improve this....?
	if c.Hide {
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	} else {
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
