package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var unpackCmd = &cobra.Command{
	Use:   "unpack",
	Short: "Unpack build artifacts from a distributed archive",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(unpackCmd).Standalone()

	rootCmd.AddCommand(unpackCmd)

	carapace.Gen(unpackCmd).PositionalCompletion(
		carapace.ActionFiles(".tgz", ".tar.gz"),
	)
}
