package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of kopsenv",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("kopsenv v0.1.0")
	},
}
