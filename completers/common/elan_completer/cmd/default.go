package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/lean"
	"github.com/spf13/cobra"
)

var defaultCmd = &cobra.Command{
	Use:   "default",
	Short: "Set the default toolchain",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(defaultCmd).Standalone()

	defaultCmd.Flags().BoolP("help", "h", false, "Prints help information")
	rootCmd.AddCommand(defaultCmd)

	carapace.Gen(defaultCmd).PositionalCompletion(
		carapace.Batch(
			lean.ActionToolchains(),
			lean.ActionToolchainSpecifiers(),
		).ToA(),
	)
}
