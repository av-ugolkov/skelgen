package cmd

import (
	"fmt"
)

var (
	Version = "v0.1.0"
)

var versionCmd = &Command{
	Name:             "version",
	ShortName:        "v",
	ShortDescription: "version for yask",
	LongDescription:  `Print the version number of yask`,
	Run: func(cmd *Command, args []string) {
		fmt.Println(Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
