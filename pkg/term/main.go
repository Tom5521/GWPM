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

	Stdout, Stderr, Stdin bool

	Bin  string
	Args []string
}

func (c *Command) Run() error {
	cmd := c.Make()
	return cmd.Run()
}

func (c *Command) Output() (string, error) {
	cmd := c.Make()
	out, err := cmd.CombinedOutput()
	return string(out), err
}
