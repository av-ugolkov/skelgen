package cmd

import (
	"fmt"
	"time"

	"github.com/av-ugolkov/gopkg/logger"
	"github.com/av-ugolkov/yask/internal/generator"

	"github.com/spf13/cobra"
)

var configPath string

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

		startTime := time.Now()
		defer func() {
			fmt.Printf("executed time: %v\n", time.Now().Sub(startTime))
		}()
		var inst map[any]any
		err := generator.GenSkeleton(configPath, inst)
		if err != nil {
			fmt.Printf("ERRORS:\n%v\n\n", err)
		}
	},
}

func Execute() {
	rootCmd.Flags().StringVarP(&configPath, "config", "c", "", "Path to config YAML(required)")
	err := rootCmd.Execute()
	if err != nil {
		logger.Errorf("%v", err)
	}
}
