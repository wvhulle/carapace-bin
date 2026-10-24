package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/lean"
	"github.com/spf13/cobra"
)

var toolchain_uninstallCmd = &cobra.Command{
	Use:     "uninstall",
	Short:   "Uninstall a toolchain",
	Aliases: []string{"remove"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(toolchain_uninstallCmd).Standalone()

	toolchain_uninstallCmd.Flags().BoolP("help", "h", false, "Prints help information")
	toolchainCmd.AddCommand(toolchain_uninstallCmd)

	carapace.Gen(toolchain_uninstallCmd).PositionalAnyCompletion(
		lean.ActionToolchains().FilterArgs(),
	)
}
