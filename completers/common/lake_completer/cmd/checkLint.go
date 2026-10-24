package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var checkLintCmd = &cobra.Command{
	Use:   "check-lint",
	Short: "Check if there is a properly configured lint driver",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(checkLintCmd).Standalone()

	rootCmd.AddCommand(checkLintCmd)
}
