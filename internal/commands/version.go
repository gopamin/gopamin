package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

const VERSION = "v1.0.10"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show current version of the tool",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Gopamin CLI Version: %v\n", VERSION)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
