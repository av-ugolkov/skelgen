package cmd

import (
	"fmt"
	"os"
	"strings"
)

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

func parseArgs(args []string) map[string][]string {
	var m map[string][]string
	for _, arg := range args {
		switch arg {
		case "-h", "--help":
			fmt.Println("Help")
		case "-v", "--version":
			fmt.Println("Version")
		case "-p", "--placeholder":
			fmt.Println("Placeholder")
		default:
			if strings.HasPrefix(arg, "-") {
				fmt.Printf("Unknown flag: %s\n", arg)
				os.Exit(1)
			}
		}
	}
	return m
}
