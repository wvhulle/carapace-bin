package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/homemanager"
	"github.com/spf13/cobra"
)

var switchCmd = &cobra.Command{
	Use:   "switch",
	Short: "Build and activate configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(switchCmd).Standalone()

	switchCmd.Flags().Bool("rollback", false, "Roll back to the configuration prior to the current configuration")
	switchCmd.Flags().StringP("specialisation", "c", "", "Activate the named specialisation")
	rootCmd.AddCommand(switchCmd)

	switchCmd.MarkFlagsMutuallyExclusive("rollback", "specialisation")

	carapace.Gen(switchCmd).FlagCompletion(carapace.ActionMap{
		"specialisation": homemanager.ActionSpecialisations(),
	})
}
