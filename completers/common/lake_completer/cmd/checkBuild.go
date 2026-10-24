package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var checkBuildCmd = &cobra.Command{
	Use:   "check-build",
	Short: "Check if any default build targets are configured",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(checkBuildCmd).Standalone()

	rootCmd.AddCommand(checkBuildCmd)
}
