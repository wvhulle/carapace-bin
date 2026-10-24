package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/lean"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a command with an environment configured for a given toolchain",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runCmd).Standalone()
	runCmd.Flags().SetInterspersed(false)

	runCmd.Flags().BoolP("help", "h", false, "Prints help information")
	runCmd.Flags().Bool("install", false, "Install the requested toolchain if needed")
	rootCmd.AddCommand(runCmd)

	carapace.Gen(runCmd).PositionalCompletion(
		carapace.Batch(
			lean.ActionToolchains(),
			lean.ActionToolchainSpecifiers(),
		).ToA(),
	)

	carapace.Gen(runCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin().Shift(1),
	)
}
