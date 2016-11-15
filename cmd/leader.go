package cmd

import (
	_ "fmt"

	"github.com/spf13/cobra"
	"github.com/w3st3ry/tekleader/tekleader"
)

func init() {
	RootCmd.AddCommand(leaderCmd)

	setLeaderFlags()
}

// leaderCmd represents the leader command
var leaderCmd = &cobra.Command{
	Use:   "leader",
	Short: "leader establish a rank betwen {EPITECH} students",
	Long: "leader establish a rank betwen {EPITECH} students" +
		" from the same promotion or city.",
	Run: func(cmd *cobra.Command, args []string) {
		tekleader.PrintLeader(tekleader.SortPromotion())
	},
}

func setLeaderFlags() {
	flags := leaderCmd.PersistentFlags()

	flags.BoolVar(&tekleader.Race, "race", false, "Enable race condition to print users")

	leaderCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
