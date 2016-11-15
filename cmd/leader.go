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

	// Custom flags
	flags.BoolVar(&tekleader.Race, "race", false, "Enable race condition to print users")
	flags.StringVar(&tekleader.Location, "location", "LYN", "Set your city (Default: Lyon)")
	flags.StringVar(&tekleader.Promo, "promotion", "tek2", "Set your promotion (Default: tek2)")

	leaderCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
