package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/lean"
	"github.com/spf13/cobra"
)

var script_docCmd = &cobra.Command{
	Use:   "doc",
	Short: "Print the docstring of a given script",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(script_docCmd).Standalone()

	scriptCmd.AddCommand(script_docCmd)

	carapace.Gen(script_docCmd).PositionalCompletion(
		inLakeWorkspace(lean.ActionLakeScripts()),
	)
}
