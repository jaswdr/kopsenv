package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of mycli",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("mycli v0.1.0")
	},
}

