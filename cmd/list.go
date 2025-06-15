package cmd

import (
	"fmt"

	"github.com/jaswdr/kopsenv/internal"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all items",
	Run: func(cmd *cobra.Command, args []string) {
		for _, release := range internal.GetReleases() {
			fmt.Printf(fmt.Sprintf("- %s\n", release.Tag))
		}
	},
}
