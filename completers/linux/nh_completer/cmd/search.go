package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Searches packages or NixOS/home-manager options via search.nixos.org, or a local SPAM database",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(searchCmd).Standalone()

	addSearchBackendFlags(searchCmd)
	searchCmd.Flags().String("default-search", "packages", "Default search mode used when no subcommand is given")
	searchCmd.Flags().BoolP("json", "j", false, "Output results as JSON")
	searchCmd.Flags().BoolP("platforms", "P", false, "Show supported platforms for each package")

	rootCmd.AddCommand(searchCmd)

	carapace.Gen(searchCmd).FlagCompletion(carapace.ActionMap{
		"default-search": carapace.ActionValuesDescribed(
			"options", "Search NixOS/home-manager options",
			"packages", "Search packages",
		),
	})
}

func addSearchBackendFlags(cmd *cobra.Command) {
	cmd.Flags().String("backend-version", "", "Backend index version to query on search.nixos.org")
	cmd.Flags().String("backend-version-fallbacks", "1", "Number of newer index versions to try when the requested version is outdated")
	cmd.Flags().StringP("channel", "c", "nixos-unstable", "Name of the channel to query")
	cmd.Flags().StringP("limit", "l", "30", "Number of search results to display")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"channel": carapace.ActionValues("nixos-unstable"),
	})
}
