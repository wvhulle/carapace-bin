package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var setupFileCmd = &cobra.Command{
	Use:    "setup-file",
	Short:  "Build the imports of a file and output its setup as JSON",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(setupFileCmd).Standalone()

	rootCmd.AddCommand(setupFileCmd)

	carapace.Gen(setupFileCmd).PositionalCompletion(
		carapace.ActionFiles(".lean"),
	)
}
