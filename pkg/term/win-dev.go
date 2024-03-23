//go:build windows
// +build windows

package term

import (
	"os/exec"
	"syscall"
)

func (c *Command) make() *exec.Cmd {
	cmd := exec.Command(c.Bin, c.Args...)
	if c.Hide {
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	}
	return cmd
}
