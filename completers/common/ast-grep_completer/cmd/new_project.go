package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/ast-grep_completer/cmd/action"
	"github.com/spf13/cobra"
)

var new_projectCmd = &cobra.Command{
	Use:   "project",
	Short: "Create an new project by scaffolding",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(new_projectCmd).Standalone()

	new_projectCmd.Flags().BoolP("help", "h", false, "Print help")
	new_projectCmd.Flags().StringP("lang", "l", "", "The language of the item to create")
	new_projectCmd.Flags().BoolP("yes", "y", false, "Accept all default options without interactive input during creation")
	newCmd.AddCommand(new_projectCmd)

	carapace.Gen(new_projectCmd).FlagCompletion(carapace.ActionMap{
		"lang": action.ActionLanguages(),
	})
}
