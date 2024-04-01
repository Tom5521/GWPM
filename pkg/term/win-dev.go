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
	}
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd
}
