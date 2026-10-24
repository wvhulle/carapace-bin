package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/lean"
	"github.com/spf13/cobra"
)

var translateConfigCmd = &cobra.Command{
	Use:   "translate-config",
	Short: "Translate a Lake configuration file into a different language",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(translateConfigCmd).Standalone()

	rootCmd.AddCommand(translateConfigCmd)

	carapace.Gen(translateConfigCmd).PositionalCompletion(
		lean.ActionLakeConfigLanguages(),
		carapace.ActionFiles(),
	)
}
