package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comparatorCmd = &cobra.Command{
	Use:   "comparator",
	Short: "Judge a solution against a challenge",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comparatorCmd).Standalone()

	comparatorCmd.Flags().String("challenge-from-export", "", "judge this export as the challenge")
	comparatorCmd.Flags().String("config", "", "JSON file describing the challenge")
	comparatorCmd.Flags().Bool("inadvisably-no-sandbox", false, "disable the built-in sandbox")
	comparatorCmd.Flags().Bool("paranoid", false, "also run all external checkers bundled with Lean")
	comparatorCmd.Flags().String("solution-from-export", "", "judge this export as the solution")
	rootCmd.AddCommand(comparatorCmd)

	carapace.Gen(comparatorCmd).FlagCompletion(carapace.ActionMap{
		"challenge-from-export": carapace.ActionFiles(),
		"config":                carapace.ActionFiles(".json"),
		"solution-from-export":  carapace.ActionFiles(),
	})
}
