//go:build !windows
// +build !windows

package term

import (
	"os"
	"os/exec"
)

// TODO: Remove this to release... or not...
func (c *Command) Make() *exec.Cmd {
	cmd := exec.Command(c.Bin, c.Args...)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd
}
