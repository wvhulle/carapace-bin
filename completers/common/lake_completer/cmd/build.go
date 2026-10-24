package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/lean"
	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build targets",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(buildCmd).Standalone()

	buildCmd.Flags().StringS("o", "o", "", "write the input-to-outputs mappings of the covered targets to a file")
	buildCmd.Flags().String("package", "", "track the targets of a different package with `-o`")
	rootCmd.AddCommand(buildCmd)

	carapace.Gen(buildCmd).FlagCompletion(carapace.ActionMap{
		"o":       carapace.ActionFiles(),
		"package": inLakeWorkspace(lean.ActionLakePackages()),
	})

	carapace.Gen(buildCmd).PositionalAnyCompletion(
		inLakeWorkspace(lean.ActionLakeTargets()),
	)
}
