package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var checkTestCmd = &cobra.Command{
	Use:   "check-test",
	Short: "Check if there is a properly configured test driver",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(checkTestCmd).Standalone()

	rootCmd.AddCommand(checkTestCmd)
}
