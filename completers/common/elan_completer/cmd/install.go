package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/lean"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:    "install",
	Short:  "Install Lean toolchain",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(installCmd).Standalone()

	installCmd.Flags().BoolP("help", "h", false, "Prints help information")
	rootCmd.AddCommand(installCmd)

	carapace.Gen(installCmd).PositionalAnyCompletion(
		lean.ActionToolchainSpecifiers().FilterArgs(),
	)
}
