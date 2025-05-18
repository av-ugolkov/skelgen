package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	cmd "github.com/av-ugolkov/yask/cmd"
)

func main() {
	args := os.Args[1:]
	ExecuteArgs(args)
}

func ExecuteArgs(args []string) {
	startTime := time.Now()
	defer func() {
		fmt.Printf("time: %v\n", time.Since(startTime))
	}()

	for ind, arg := range args {
		switch arg {
		case "-h", "--help":
			cmd.RunHelp()
		case "-v", "--version":
			cmd.RunVersion()
		case "-c", "--config":
			configPath := args[ind+1]
			for _, arg := range args[ind+2:] {
				switch arg {
				case "-p", "--placeholder":
					fmt.Println("Placeholder")
				}
			}
			cmd.RunGenSkel(configPath, args[ind+2:])
		default:
			if strings.HasPrefix(arg, "-") {
				fmt.Printf("Unknown flag: %s\n", arg)
				os.Exit(1)
			}
		}
	}
}
