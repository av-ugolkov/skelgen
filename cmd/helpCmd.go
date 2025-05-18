package cmd

import (
	"fmt"
)

const help string = `Generate project structure from YAML

Usage:
  yask [flags]
  yask [command]

Available Commands:
  help        Help about any command
  version     Print the version number of yask

Flags:
  -c, --config string       Path to config YAML(required)
  -d, --dynamic key=value   Dynamics arguments in format [key] [value]
  -h, --help                help for yask

Use "yask [command] --help" for more information about a command.
`

var helpCmd = &Command{
	Name:             "help",
	ShortName:        "h",
	ShortDescription: "help for yask",
	LongDescription:  "Help about any command",
	Run: func(cmd *Command, args []string) {
		fmt.Print(help)
	},
}

func init() {
	rootCmd.AddCommand(helpCmd)
}

func RunHelp() {
	helpCmd.RunCommand(nil)
}
