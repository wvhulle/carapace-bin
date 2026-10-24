package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var new_helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Print this message or the help of the given subcommand(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(new_helpCmd).Standalone()

	newCmd.AddCommand(new_helpCmd)

	carapace.Gen(new_helpCmd).PositionalAnyCompletion(
		carapace.ActionCommands(newCmd),
	)
}
