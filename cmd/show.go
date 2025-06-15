package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show [item]",
	Short: "Show details of an item",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		item := args[0]
		fmt.Printf("Details for %s:\n", item)
		// Simulated data
		fmt.Println("Name:", item)
		fmt.Println("Size: 1.2MB")
		fmt.Println("Date: 2025-06-15")
	},
}

