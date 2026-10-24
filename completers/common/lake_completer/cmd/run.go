package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/lean"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Shorthand for `lake script run`",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runCmd).Standalone()
	runCmd.Flags().SetInterspersed(false)

	rootCmd.AddCommand(runCmd)

	carapace.Gen(runCmd).PositionalCompletion(
		inLakeWorkspace(lean.ActionLakeScripts()),
	)

	carapace.Gen(runCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
