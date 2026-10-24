package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/ast-grep_completer/cmd/action"
	"github.com/spf13/cobra"
)

var new_testCmd = &cobra.Command{
	Use:   "test",
	Short: "Create a new test case",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(new_testCmd).Standalone()

	new_testCmd.Flags().BoolP("help", "h", false, "Print help")
	new_testCmd.Flags().StringP("lang", "l", "", "The language of the item to create")
	new_testCmd.Flags().BoolP("yes", "y", false, "Accept all default options without interactive input during creation")
	newCmd.AddCommand(new_testCmd)

	carapace.Gen(new_testCmd).FlagCompletion(carapace.ActionMap{
		"lang": action.ActionLanguages(),
	})
}
