package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var toolchain_helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Prints this message or the help of the given subcommand(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(toolchain_helpCmd).Standalone()

	toolchainCmd.AddCommand(toolchain_helpCmd)

	carapace.Gen(toolchain_helpCmd).PositionalAnyCompletion(
		carapace.ActionCommands(toolchainCmd),
	)
}
