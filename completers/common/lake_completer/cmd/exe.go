package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/lean"
	"github.com/spf13/cobra"
)

var exeCmd = &cobra.Command{
	Use:     "exe",
	Short:   "Build an executable target and run it in Lake's environment",
	Aliases: []string{"exec"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(exeCmd).Standalone()
	exeCmd.Flags().SetInterspersed(false)

	rootCmd.AddCommand(exeCmd)

	carapace.Gen(exeCmd).PositionalCompletion(
		inLakeWorkspace(lean.ActionLakeExecutables()),
	)

	carapace.Gen(exeCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
