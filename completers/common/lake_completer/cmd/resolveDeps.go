package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resolveDepsCmd = &cobra.Command{
	Use:    "resolve-deps",
	Short:  "Resolve dependencies without building",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resolveDepsCmd).Standalone()

	rootCmd.AddCommand(resolveDepsCmd)
}
