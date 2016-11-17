package cmd

import (
	"github.com/spf13/cobra"
	"github.com/w3st3ry/tekleader/tekleader"
)

func init() {
	RootCmd.AddCommand(statusCmd)
}

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Give intranet status access in continue",
	Run: func(cmd *cobra.Command, args []string) {
		tekleader.PrintStatus(true)
	},
}
