package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cache_stageCmd = &cobra.Command{
	Use:   "stage",
	Short: "Copy build outputs from the cache to a staging directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cache_stageCmd).Standalone()

	cache_stageCmd.Flags().Bool("force-overwrite", false, "overwrite artifacts in the staging directory")
	cacheCmd.AddCommand(cache_stageCmd)

	carapace.Gen(cache_stageCmd).PositionalCompletion(
		carapace.ActionFiles(),
		carapace.ActionDirectories(),
	)
}
