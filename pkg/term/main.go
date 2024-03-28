package term

import msg "github.com/Tom5521/GoNotes/pkg/messages"

func NewCommand(bin string, args ...string) *Command {
	cmd := &Command{
		Bin:  bin,
		Args: args,
	}
	return cmd
}

type Command struct {
	Hide bool

	Stdout, Stderr, Stdin bool

	Bin  string
	Args []string
}

func (c *Command) Run() error {
	cmd := c.Make()
	msg.Info("Running ", cmd)
	return cmd.Run()
}

func (c *Command) Output() (string, error) {
	cmd := c.Make()
	out, err := cmd.CombinedOutput()
	msg.Info("Running ", cmd)
	return string(out), err
}
