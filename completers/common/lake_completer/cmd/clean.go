package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/lean"
	"github.com/spf13/cobra"
)

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Remove build outputs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanCmd).Standalone()

	rootCmd.AddCommand(cleanCmd)

	carapace.Gen(cleanCmd).PositionalAnyCompletion(
		inLakeWorkspace(lean.ActionLakePackages()).FilterArgs(),
	)
}
