package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var leanCmd = &cobra.Command{
	Use:   "lean",
	Short: "Elaborate a Lean file in the context of the Lake workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(leanCmd).Standalone()

	rootCmd.AddCommand(leanCmd)

	carapace.Gen(leanCmd).PositionalCompletion(
		carapace.ActionFiles(".lean"),
	)

	carapace.Gen(leanCmd).DashAnyCompletion(
		bridge.ActionCarapaceBin("lean"),
	)
}
