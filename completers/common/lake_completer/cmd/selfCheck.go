package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var selfCheckCmd = &cobra.Command{
	Use:    "self-check",
	Short:  "Output the detected installations and verify the Lean version",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(selfCheckCmd).Standalone()

	rootCmd.AddCommand(selfCheckCmd)
}
