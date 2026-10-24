package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var versionTagsCmd = &cobra.Command{
	Use:    "version-tags",
	Short:  "Print the version tags of the package",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(versionTagsCmd).Standalone()

	rootCmd.AddCommand(versionTagsCmd)
}
