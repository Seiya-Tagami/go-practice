package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "get the version",
	Long:  `get the version`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("v1.0.1")
	},
}
