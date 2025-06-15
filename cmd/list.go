package cmd

import (
	"fmt"

	"github.com/jaswdr/kopsenv/internal"

	"github.com/spf13/cobra"
)

var listAll bool

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available versions",
	Run: func(cmd *cobra.Command, args []string) {
		allReleases := internal.GetReleases()
		if !listAll {
			allReleases = allReleases[:10]
		}
		for _, release := range allReleases {
			fmt.Println(release)
		}
	},
}

func init() {
	listCmd.Flags().BoolVar(&listAll, "all", false, "List all flags instead of latest 10")
}
