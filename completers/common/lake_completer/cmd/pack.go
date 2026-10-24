package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var packCmd = &cobra.Command{
	Use:   "pack",
	Short: "Pack build artifacts into an archive for distribution",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(packCmd).Standalone()

	rootCmd.AddCommand(packCmd)

	carapace.Gen(packCmd).PositionalCompletion(
		carapace.ActionFiles(".tgz", ".tar.gz"),
	)
}
