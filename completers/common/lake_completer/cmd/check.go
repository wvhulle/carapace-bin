package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check this project against external checker(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(checkCmd).Standalone()

	checkCmd.Flags().String("from-export", "", "check this export, instead of building and exporting the project")
	checkCmd.Flags().Bool("inadvisably-no-sandbox", false, "disable the built-in sandbox")
	checkCmd.Flags().Bool("paranoid", false, "also run all external checkers bundled with Lean")
	rootCmd.AddCommand(checkCmd)

	carapace.Gen(checkCmd).FlagCompletion(carapace.ActionMap{
		"from-export": carapace.ActionFiles(),
	})
}
