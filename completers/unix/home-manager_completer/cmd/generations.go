package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var generationsCmd = &cobra.Command{
	Use:   "generations",
	Short: "List all home environment generations",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(generationsCmd).Standalone()

	rootCmd.AddCommand(generationsCmd)
}
