package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var packagesCmd = &cobra.Command{
	Use:   "packages",
	Short: "List all packages installed in home-manager-path",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(packagesCmd).Standalone()

	rootCmd.AddCommand(packagesCmd)
}
