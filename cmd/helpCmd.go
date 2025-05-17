package cmd

import (
	"fmt"
	"os"
)

var helpCmd = &Command{
	Name:             "help",
	ShortName:        "h",
	ShortDescription: "help for yask",
	LongDescription:  "Help about any command",
	Run: func(cmd *Command, args []string) {
		fmt.Println("Unimplemented")
		os.Exit(1)
	},
}

func init() {
	rootCmd.AddCommand(helpCmd)
}
