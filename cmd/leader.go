package cmd

import (
	str "strings"

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
	if tekleader.Location == "*" {
		fmtArgs := make([]string, len(cities))

		// Iterating on all cities
		i := 0
		for _, city := range cities {
			fmtArgs[i] = "FR/" + city
			i++
		}

		tekleader.Location = str.Join(fmtArgs, "|")
	} else {
		args := splitArgs(tekleader.Location)
		fmtArgs := make([]string, len(args))

		// Iterating on all given cities
		for key, city := range cities {
			for i, arg := range args {
				if key == arg {
					fmtArgs[i] = "FR/" + city
				}
			}
		}

		tekleader.Location = str.Join(fmtArgs, "|")
	}
}

func formatFind() {
	if len(tekleader.Find) == 0 {
		return
	}

	args := splitArgs(tekleader.Find)
	fmtArgs := make([]string, len(args))
	for i, login := range args {
		if !str.Contains(login, "@") {
			fmtArgs[i] = login + "@epitech.eu"
		} else {
			fmtArgs[i] = login
		}
	}
	tekleader.Find = str.Join(fmtArgs, ",")
}

func splitArgs(s string) []string {
	s = str.Replace(str.ToLower(s), " ", "", -1)
	return str.Split(s, ",")
}
