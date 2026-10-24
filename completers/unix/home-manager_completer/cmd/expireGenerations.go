package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var expireGenerationsCmd = &cobra.Command{
	Use:   "expire-generations",
	Short: "Remove generations older than TIMESTAMP",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(expireGenerationsCmd).Standalone()

	rootCmd.AddCommand(expireGenerationsCmd)

	carapace.Gen(expireGenerationsCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"-1 days", "Older than one day",
			"-7 days", "Older than one week",
			"-30 days", "Older than one month",
			"-1 year", "Older than one year",
		),
	)
}
