package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/ast-grep_completer/cmd/action"
	"github.com/spf13/cobra"
)

var new_ruleCmd = &cobra.Command{
	Use:   "rule",
	Short: "Create a new rule",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(new_ruleCmd).Standalone()

	new_ruleCmd.Flags().BoolP("help", "h", false, "Print help")
	new_ruleCmd.Flags().StringP("lang", "l", "", "The language of the item to create")
	new_ruleCmd.Flags().BoolP("yes", "y", false, "Accept all default options without interactive input during creation")
	newCmd.AddCommand(new_ruleCmd)

	carapace.Gen(new_ruleCmd).FlagCompletion(carapace.ActionMap{
		"lang": action.ActionLanguages(),
	})
}
