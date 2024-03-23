package term

func NewCommand(bin string, args ...string) *Command {
	cmd := &Command{
		Bin:  bin,
		Args: args,
	}
	return cmd
}

type Command struct {
	Hide bool

	Bin  string
	Args []string
}

func (c *Command) Run() error {
	cmd := c.make()
	return cmd.Run()
}

func (c *Command) Output() (string, error) {
	cmd := c.make()
	out, err := cmd.Output()
	return string(out), err
}
