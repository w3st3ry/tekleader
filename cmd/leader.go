package cmd

import (
	"strings"

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
		setTekCourse()
		formatCities()
		formatFind()
		tekleader.PrintLeader(tekleader.SortPromotion())
	},
}

func setLeaderFlags() {
	flags := leaderCmd.PersistentFlags()

	// Custom flags
	flags.BoolVar(&tekleader.Race, "race", false, "Enable race condition to print users")
	flags.StringVar(&tekleader.Location, "location", "lyon", "Set your city")
	flags.StringVar(&tekleader.Promo, "promotion", "tek2", "Set your promotion")
	flags.StringVar(&tekleader.Find, "find", "", "The student you want to find by login")

	leaderCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// Change course for tek3+
func setTekCourse() {
	p := tekleader.Promo
	if p == "tek4" || p == "tek5" {
		tekleader.Course = "master"
	} else {
		tekleader.Course = "bachelor"
	}
}

func formatCities() {
	f := strings.ToLower(tekleader.Location)
	cities := map[string]string{
		"lyon":        "LYN",
		"paris":       "PAR",
		"bordeaux":    "BDX",
		"marseille":   "MAR",
		"lille":       "LIL",
		"montpellier": "MPL",
		"nancy":       "NCY",
		"nantes":      "NAN",
		"nice":        "NCE",
		"rennes":      "REN",
		"strasbourg":  "STG",
		"toulouse":    "TLS",
	}
	for key, city := range cities {
		if key == f {
			tekleader.Location = city
			break
		}
	}
}

func formatFind() {
	f := strings.ToLower(tekleader.Find)
	f = strings.Trim(f, " ")
	f = strings.Replace(f, " ", ".", -1)
	if !strings.Contains(f, "@epitech.eu") {
		tekleader.Find = f + "@epitech.eu"
	}
}
