package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/homemanager"
	"github.com/spf13/cobra"
)

var removeGenerationsCmd = &cobra.Command{
	Use:   "remove-generations",
	Short: "Remove indicated generations",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(removeGenerationsCmd).Standalone()

	rootCmd.AddCommand(removeGenerationsCmd)

	carapace.Gen(removeGenerationsCmd).PositionalAnyCompletion(
		homemanager.ActionGenerations().FilterArgs(),
	)
}
