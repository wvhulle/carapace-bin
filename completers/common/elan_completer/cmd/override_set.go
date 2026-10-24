package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/lean"
	"github.com/spf13/cobra"
)

var override_setCmd = &cobra.Command{
	Use:     "set",
	Short:   "Set the override toolchain for a directory",
	Aliases: []string{"add"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(override_setCmd).Standalone()

	override_setCmd.Flags().BoolP("help", "h", false, "Prints help information")
	overrideCmd.AddCommand(override_setCmd)

	carapace.Gen(override_setCmd).PositionalCompletion(
		carapace.Batch(
			lean.ActionToolchains(),
			lean.ActionToolchainSpecifiers(),
		).ToA(),
	)
}
