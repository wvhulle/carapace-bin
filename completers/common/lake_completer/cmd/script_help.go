package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var script_helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Print help of a script command",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(script_helpCmd).Standalone()

	scriptCmd.AddCommand(script_helpCmd)

	carapace.Gen(script_helpCmd).PositionalCompletion(
		carapace.ActionCommands(scriptCmd),
	)
}
