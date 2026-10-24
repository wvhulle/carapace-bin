package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var search_prsCmd = &cobra.Command{
	Use:   "prs",
	Short: "Search Nixpkgs pull requests and branch reachability",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(search_prsCmd).Standalone()

	search_prsCmd.Flags().StringP("days", "d", "15", "Search GitHub results updated in the last n days")
	search_prsCmd.Flags().BoolP("json", "j", false, "Output results as JSON")

	searchCmd.AddCommand(search_prsCmd)
}
