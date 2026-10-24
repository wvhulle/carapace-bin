package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var search_offlineCmd = &cobra.Command{
	Use:   "offline",
	Short: "Search local SPAM database(s) without network access",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(search_offlineCmd).Standalone()

	search_offlineCmd.Flags().StringArrayP("db", "D", nil, "Path to a SPAM database file")
	search_offlineCmd.Flags().BoolP("json", "j", false, "Output results as JSON")
	search_offlineCmd.Flags().StringP("limit", "l", "30", "Number of search results to display")

	search_offlineCmd.MarkFlagRequired("db")

	searchCmd.AddCommand(search_offlineCmd)

	carapace.Gen(search_offlineCmd).FlagCompletion(carapace.ActionMap{
		"db": carapace.ActionFiles(),
	})
}
