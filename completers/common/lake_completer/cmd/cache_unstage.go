package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cache_unstageCmd = &cobra.Command{
	Use:   "unstage",
	Short: "Cache build outputs from a staging directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cache_unstageCmd).Standalone()

	cache_unstageCmd.Flags().Bool("force-overwrite", false, "overwrite mappings and artifacts already in the cache")
	cacheCmd.AddCommand(cache_unstageCmd)

	carapace.Gen(cache_unstageCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
