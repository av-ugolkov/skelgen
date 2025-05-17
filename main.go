package main

import (
	"os"
)

import "github.com/av-ugolkov/yask/cmd"

func main() {
	args := os.Args[1:]
	cmd.ExecuteArgs(args)
}

/*
Generate project structure from YAML

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
*/
