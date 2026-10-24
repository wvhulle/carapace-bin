package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/lean"
	"github.com/spf13/cobra"
)

var queryKindCmd = &cobra.Command{
	Use:    "query-kind",
	Short:  "Output the kind of targets",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(queryKindCmd).Standalone()

	rootCmd.AddCommand(queryKindCmd)

	carapace.Gen(queryKindCmd).PositionalAnyCompletion(
		inLakeWorkspace(lean.ActionLakeTargets()),
	)
}
