package term

import "os/exec"

func (c *Command) make() *exec.Cmd {
	cmd := exec.Command(c.Bin, c.Args...)
	return cmd
}
