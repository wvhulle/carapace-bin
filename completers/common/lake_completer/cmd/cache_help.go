package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cache_helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Print help of a cache command",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cache_helpCmd).Standalone()

	cacheCmd.AddCommand(cache_helpCmd)

	carapace.Gen(cache_helpCmd).PositionalCompletion(
		carapace.ActionCommands(cacheCmd),
	)
}
