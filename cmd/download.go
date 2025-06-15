package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var downloadCmd = &cobra.Command{
	Use:   "download [version]",
	Short: "Downloads a particular kops version",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		item := args[0]
		fmt.Printf("Downloading %s...\n", item)
		// Simulated delay or logic here
		fmt.Println("Download complete.")
	},
}
