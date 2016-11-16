package cmd

import (
	"errors"
	"os"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/w3st3ry/tekleader/tekleader"
)

func init() {
	// Set logger format
	formatter := new(log.TextFormatter)
	formatter.TimestampFormat = time.Stamp
	formatter.FullTimestamp = true
	log.SetFormatter(formatter)
	log.SetOutput(os.Stdout)

	setRootFlags()
}

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use: "Tekleader",
	Short: "Tekleader makes it possible to establish" +
		" a rank between {EPITECH} students, and much more.",
	Long: "",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// Check timeout value
		if tekleader.Timeout < 1 {
			log.Fatal(errors.New("Timeout must be > 1"))
		}

		// Skip auth for defined commands
		if skipAuth(cmd.Name()) {
			return
		}

		// Check if intranet is alive before call any endpoints
		tekleader.PrintStatus(false)

		// Use autkey in cfgfile or env if exist
		tekleader.AuthKey = viper.GetString("tek_authkey")

		// Try to log in with defined authkey, abort if unsuccessful
		err := tekleader.Auth()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func setRootFlags() {
	flags := RootCmd.PersistentFlags()

	// Global flags
	flags.StringVar(&tekleader.AuthKey, "auth-key", "", "config file (default is $HOME/.tekleader.yml)")
	flags.IntVar(&tekleader.Timeout, "timeout", 2, "Timeout must be > 1s")

	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// Bind cobra flags to viper flags
	viper.BindPFlag("tek_authkey", RootCmd.PersistentFlags().Lookup("auth-key"))
	viper.BindPFlag("tek_timeout", RootCmd.PersistentFlags().Lookup("timeout"))
}

// skipAuth match commands who won't need authentication
func skipAuth(command string) bool {
	deny := [2]string{"version", "status"}
	for _, key := range deny {
		if key == command {
			return true
		}
	}
	return false
}
