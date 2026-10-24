package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var search_optionsCmd = &cobra.Command{
	Use:   "options",
	Short: "Search NixOS/home-manager options via search.nixos.org",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(search_optionsCmd).Standalone()

	addSearchBackendFlags(search_optionsCmd)
	search_optionsCmd.Flags().BoolP("json", "j", false, "Output results as JSON")
	search_optionsCmd.Flags().String("scope", "", "Options scope")

	search_optionsCmd.Flag("scope").NoOptDefVal = "all"

	searchCmd.AddCommand(search_optionsCmd)

	carapace.Gen(search_optionsCmd).FlagCompletion(carapace.ActionMap{
		"scope": carapace.ActionValuesDescribed(
			"all", "Search all options",
			"home-manager", "Search home-manager options",
			"nix-darwin", "Search nix-darwin options",
			"nixpkgs", "Search NixOS options and modular services",
		),
	})
}
