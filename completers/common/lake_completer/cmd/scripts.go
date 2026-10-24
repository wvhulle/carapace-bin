package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var scriptsCmd = &cobra.Command{
	Use:   "scripts",
	Short: "Shorthand for `lake script list`",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(scriptsCmd).Standalone()

	rootCmd.AddCommand(scriptsCmd)
}
