package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/lean"
	"github.com/spf13/cobra"
)

var toolchain_installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install a given toolchain",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(toolchain_installCmd).Standalone()

	toolchain_installCmd.Flags().BoolP("help", "h", false, "Prints help information")
	toolchainCmd.AddCommand(toolchain_installCmd)

	carapace.Gen(toolchain_installCmd).PositionalAnyCompletion(
		lean.ActionToolchainSpecifiers().FilterArgs(),
	)
}
