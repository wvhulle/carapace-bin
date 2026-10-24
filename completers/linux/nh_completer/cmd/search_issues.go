package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var search_issuesCmd = &cobra.Command{
	Use:   "issues",
	Short: "Search Nixpkgs issues, excluding pull requests",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(search_issuesCmd).Standalone()

	search_issuesCmd.Flags().StringP("days", "d", "15", "Search GitHub results updated in the last n days")
	search_issuesCmd.Flags().BoolP("json", "j", false, "Output results as JSON")

	searchCmd.AddCommand(search_issuesCmd)
}
