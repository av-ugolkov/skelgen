package cmd

import (
	"fmt"
	"strings"

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
	rootCmd.RunCommand(args)
}
