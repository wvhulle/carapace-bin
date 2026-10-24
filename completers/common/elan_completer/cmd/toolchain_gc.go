package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var toolchain_gcCmd = &cobra.Command{
	Use:   "gc",
	Short: "Garbage-collect toolchains not used by any known project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(toolchain_gcCmd).Standalone()

	toolchain_gcCmd.Flags().Bool("delete", false, "Delete collected toolchains instead of only reporting them")
	toolchain_gcCmd.Flags().BoolP("help", "h", false, "Prints help information")
	toolchain_gcCmd.Flags().Bool("json", false, "Format output as JSON")
	toolchainCmd.AddCommand(toolchain_gcCmd)
}
