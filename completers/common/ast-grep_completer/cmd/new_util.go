package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/ast-grep_completer/cmd/action"
	"github.com/spf13/cobra"
)

var new_utilCmd = &cobra.Command{
	Use:   "util",
	Short: "Create a new global utility rule",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(new_utilCmd).Standalone()

	new_utilCmd.Flags().BoolP("help", "h", false, "Print help")
	new_utilCmd.Flags().StringP("lang", "l", "", "The language of the item to create")
	new_utilCmd.Flags().BoolP("yes", "y", false, "Accept all default options without interactive input during creation")
	newCmd.AddCommand(new_utilCmd)

	carapace.Gen(new_utilCmd).FlagCompletion(carapace.ActionMap{
		"lang": action.ActionLanguages(),
	})
}
