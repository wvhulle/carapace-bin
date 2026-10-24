package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/lean"
	"github.com/spf13/cobra"
)

var whichCmd = &cobra.Command{
	Use:   "which",
	Short: "Display which binary will be run for a given command",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(whichCmd).Standalone()

	whichCmd.Flags().BoolP("help", "h", false, "Prints help information")
	rootCmd.AddCommand(whichCmd)

	carapace.Gen(whichCmd).PositionalCompletion(
		lean.ActionToolchainBinaries(),
	)
}
