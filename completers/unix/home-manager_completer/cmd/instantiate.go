package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var instantiateCmd = &cobra.Command{
	Use:   "instantiate",
	Short: "Instantiate the configuration and print the resulting derivation",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(instantiateCmd).Standalone()

	rootCmd.AddCommand(instantiateCmd)
}
