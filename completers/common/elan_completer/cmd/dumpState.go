package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dumpStateCmd = &cobra.Command{
	Use:    "dump-state",
	Short:  "Dump the elan state as JSON",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dumpStateCmd).Standalone()

	dumpStateCmd.Flags().BoolP("help", "h", false, "Prints help information")
	dumpStateCmd.Flags().Bool("no-net", false, "Make network operations for resolving channels fail immediately")
	rootCmd.AddCommand(dumpStateCmd)
}
