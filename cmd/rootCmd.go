package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/av-ugolkov/gopkg/logger"
	"github.com/av-ugolkov/yask/internal/generator"

	"github.com/spf13/cobra"
)

var configPath string
var kv kvPairs

var rootCmd = &cobra.Command{
	Use:   "yask",
	Short: "Generate project structure from YAML",
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
	Run: func(cmd *cobra.Command, args []string) {
		if configPath == "" {
			err := cmd.Help()
			if err != nil {
				logger.Errorf("%v", err)
			}
			return
		}
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

func Execute() {
	startTime := time.Now()
	defer func() {
		fmt.Printf("time: %v\n", time.Since(startTime))
	}()
	rootCmd.Flags().StringVarP(&configPath, "config", "c", "", "Path to config YAML(required)")
	rootCmd.Flags().VarP(&kv, "dynamic", "d", "Dynamics arguments in format [key] [value]")
	err := rootCmd.Execute()
	if err != nil {
		logger.Errorf("%v", err)
	}
}
