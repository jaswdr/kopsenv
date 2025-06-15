package cmd

import (
	"fmt"

	"github.com/jaswdr/kopsenv/internal"
	"github.com/spf13/cobra"
)

var useCmd = &cobra.Command{
	Use:   "use [version]",
	Short: "Make use a particular kops version",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		version := args[0]
		if !internal.IsVersionAvailable(version) {
			fmt.Println("Version is not available, please download it first")
		}

		internal.LinkVersion(version)
	},
}
