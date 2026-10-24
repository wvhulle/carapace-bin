package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/ast-grep_completer/cmd/action"
	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create new ast-grep project or items like rules/tests",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(newCmd).Standalone()

	newCmd.Flags().BoolP("help", "h", false, "Print help")
	newCmd.Flags().StringP("lang", "l", "", "The language of the item to create")
	newCmd.Flags().BoolP("yes", "y", false, "Accept all default options without interactive input during creation")
	rootCmd.AddCommand(newCmd)

	carapace.Gen(newCmd).FlagCompletion(carapace.ActionMap{
		"lang": action.ActionLanguages(),
	})
}
