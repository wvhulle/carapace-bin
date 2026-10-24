package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/lean"
	"github.com/spf13/cobra"
)

var queryCmd = &cobra.Command{
	Use:   "query",
	Short: "Build targets and output results",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(queryCmd).Standalone()

	rootCmd.AddCommand(queryCmd)

	carapace.Gen(queryCmd).PositionalAnyCompletion(
		inLakeWorkspace(lean.ActionLakeTargets()),
	)
}
