package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	Version   = "dev"
	Commit    = "none"
	BuildDate = "unknown"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of yask",
	Long:  `All software has versions. This is yask's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("yask:", Version)
		fmt.Println("yask commit:", Commit)
		fmt.Println("yask build date:", BuildDate)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
