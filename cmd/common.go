package cmd

type Command struct {
	Name             string
	ShortName        string
	ShortDescription string
	LongDescription  string
	Run              func(cmd *Command, args []string)

	childCommands []*Command
}

func (c *Command) RunCommand(args []string) {
	c.Run(c, args)
}

func (c *Command) AddCommand(cmd *Command) {
	cmd.childCommands = append(cmd.childCommands, cmd)
}
