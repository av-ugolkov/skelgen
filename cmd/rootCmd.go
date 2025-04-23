package cmd

import (
	"os"

	"github.com/av-ugolkov/gopkg/logger"
	"github.com/av-ugolkov/yask/internal"

	"github.com/spf13/cobra"
)

var configPath string

var rootCmd = &cobra.Command{
	Use:   "yask",
	Short: "Generate project structure from YAML",
	Run: func(cmd *cobra.Command, args []string) {
		if configPath == "" {
			logger.Errorf("path is required")
			os.Exit(1)
		}

		var inst map[any]any
		err := internal.GenSkeleton(configPath, inst)
		if err != nil {
			logger.Errorf("%v", err)
			os.Exit(1)
		}
	},
}

func Execute() {
	rootCmd.Flags().StringVarP(&configPath, "config", "c", "", "Path to config YAML")
	cobra.CheckErr(rootCmd.Execute())
}
