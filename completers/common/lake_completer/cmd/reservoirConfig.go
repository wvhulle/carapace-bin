package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var reservoirConfigCmd = &cobra.Command{
	Use:    "reservoir-config",
	Short:  "Output the Reservoir configuration of the package as JSON",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(reservoirConfigCmd).Standalone()

	rootCmd.AddCommand(reservoirConfigCmd)
}
