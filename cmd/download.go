package cmd

import (
	"fmt"

	"github.com/jaswdr/kopsenv/internal"
	"github.com/spf13/cobra"
)

var (
	downloadLatest       bool
	downloadLatestStable bool
)

var downloadCmd = &cobra.Command{
	Use:   "download [version]",
	Short: "Downloads a particular kops version",
	Args:  cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			if !downloadLatest && !downloadLatestStable {
				fmt.Println("Either a version, explicit --latest or --latest-stable should be used")
			}
		}

		var version string
		if downloadLatest || downloadLatestStable {
			releases := internal.GetReleases()
			if downloadLatest {
				version = releases[0].Tag
			}

			for _, release := range releases {
				if !release.IsAlpha && !release.IsBeta {
					version = release.Tag
					break
				}
			}
		} else {
			version = args[0]
		}

		internal.Download(version)
	},
}

func init() {
	downloadCmd.Flags().BoolVar(&downloadLatest, "latest", false, "Download the latest version")
	downloadCmd.Flags().BoolVar(&downloadLatestStable, "latest-stable", false, "Download the latest stable version")
}
