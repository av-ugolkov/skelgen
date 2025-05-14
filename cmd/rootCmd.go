package cmd

import (
	"fmt"
	"time"

	"github.com/av-ugolkov/gopkg/logger"
	"github.com/av-ugolkov/yask/internal/generator"

	"github.com/spf13/cobra"
)

var configPath string
var dynamic []string

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
		dynamicM := make(map[string]string, len(dynamic)/2)
		for i := 0; i < len(dynamic); i = i + 2 {
			dynamicM[dynamic[i]] = dynamic[i+1]
		}

		var inst map[any]any
		err := generator.GenSkeleton(configPath, inst, dynamicM)
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
	rootCmd.Flags().StringSliceVarP(&dynamic, "dynamic", "d", nil, "Path to config YAML(required)")
	err := rootCmd.Execute()
	if err != nil {
		logger.Errorf("%v", err)
	}
}
