package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/lean"
	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a Lean package in a new directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(newCmd).Standalone()

	rootCmd.AddCommand(newCmd)

	carapace.Gen(newCmd).PositionalCompletion(
		carapace.ActionValues(),
		lean.ActionLakeTemplates(),
	)
}
