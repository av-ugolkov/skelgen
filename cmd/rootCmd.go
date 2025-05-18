package cmd

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/av-ugolkov/yask/internal/generator"
)

var configPath string
var kv kvPairs

var rootCmd = &Command{
	Name:             "yask",
	ShortDescription: "Generate project structure from YAML",
	Run: func(cmd *Command, args []string) {
		dynamic := make(map[string]string, len(kv))
		for k, v := range kv {
			dynamic[strings.TrimSpace(k)] = strings.TrimSpace(v)
		}

		var inst map[any]any
		err := generator.GenSkeleton(configPath, inst, dynamic)
		if err != nil {
			fmt.Printf("ERRORS:\n%v\n\n", err)
		}
	},
}

// func Execute() {
// 	startTime := time.Now()
// 	defer func() {
// 		fmt.Printf("time: %v\n", time.Since(startTime))
// 	}()
// 	rootCmd.Flags().StringVarP(&configPath, "config", "c", "", "Path to config YAML(required)")
// 	rootCmd.Flags().VarP(&kv, "dynamic", "d", "Dynamics arguments in format [key] [value]")
// 	err := rootCmd.Execute()
// 	if err != nil {
// 		logger.Errorf("%v", err)
// 	}
// }

func ExecuteArgs(args []string) {
	startTime := time.Now()
	defer func() {
		fmt.Printf("time: %v\n", time.Since(startTime))
	}()

	for ind, arg := range args {
		switch arg {
		case "-h", "--help":
			helpCmd.RunCommand(nil)
		case "-v", "--version":
			versionCmd.RunCommand(nil)
		case "-c", "--config":
			configPath = args[ind+1]
			for _, arg := range args[ind+2:] {
				switch arg {
				case "-p", "--placeholder":
					fmt.Println("Placeholder")
				}
			}
			rootCmd.RunCommand(args[ind+2:])
		default:
			if strings.HasPrefix(arg, "-") {
				fmt.Printf("Unknown flag: %s\n", arg)
				os.Exit(1)
			}
		}
	}
}
