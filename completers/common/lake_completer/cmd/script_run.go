package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/lean"
	"github.com/spf13/cobra"
)

var script_runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a script",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(script_runCmd).Standalone()
	script_runCmd.Flags().SetInterspersed(false)

	scriptCmd.AddCommand(script_runCmd)

	carapace.Gen(script_runCmd).PositionalCompletion(
		inLakeWorkspace(lean.ActionLakeScripts()),
	)

	carapace.Gen(script_runCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
