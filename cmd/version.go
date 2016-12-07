package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/google/go-github/github"
	"github.com/spf13/cobra"
	"github.com/w3st3ry/tekleader/tekleader"
)

const (
	cVersion    = "v" + tekleader.Version
	repOwner    = "w3st3ry"
	projectName = "tekleader"
)

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the tekleader version and any available update",
	Long: "Display information about current tekleader version\n" +
		"and check if a new release is available on GitHub.",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := github.NewClient(nil)

		color.White("Version information:\n")
		current, _, err := client.Repositories.GetReleaseByTag(repOwner, projectName, cVersion)
		printRelease(current, err)
		if err != nil {
			return nil
		}

		releases, _, err := client.Repositories.ListReleases(repOwner, projectName, &github.ListOptions{PerPage: 1})
		latest := releases[len(releases)-1]

		if *current.TagName == *latest.TagName {
			color.Green("Your version is up to date.")
		} else {
			color.Yellow("A new version is available:\n")
			printRelease(latest, err)
			color.Yellow("Visit %s for more information about tekleader", *latest.HTMLURL)
		}

		return err
	},
}

func printRelease(release *github.RepositoryRelease, err error) {
	if err != nil {
		color.Red("Your version is not recognized or invalid! Check the error(s) below.\n%v", err)
		return
	}

	print := fmt.Printf
	print("* Release name: %s\n", *release.Name)
	print("* Related git tag: %s\n", *release.TagName)
	print("* Stable release: %t\n", !*release.Prerelease)

	if release.PublishedAt != nil && !release.PublishedAt.IsZero() {
		print("* Release date: %s\n", *release.PublishedAt)
	}
}
