package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var optionCmd = &cobra.Command{
	Use:   "option",
	Short: "Inspect configuration option named OPTION.NAME",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(optionCmd).Standalone()

	optionCmd.Flags().Bool("recursive", false, "Print all the values at or below the option name recursively")
	rootCmd.AddCommand(optionCmd)
}
