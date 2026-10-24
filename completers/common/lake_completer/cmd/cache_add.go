package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/lean"
	"github.com/spf13/cobra"
)

var cache_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add input-to-output mappings to the Lake cache",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cache_addCmd).Standalone()

	cache_addCmd.Flags().Bool("no-overwrite", false, "do not overwrite existing mappings")
	cache_addCmd.Flags().String("package", "", "add mappings to set package")
	cache_addCmd.Flags().String("repo", "", "for Reservoir, a GitHub repository scope")
	cache_addCmd.Flags().String("scope", "", "the prefix of artifacts within the service")
	cache_addCmd.Flags().String("service", "", "cache service to fetch from on demand")
	cacheCmd.AddCommand(cache_addCmd)

	carapace.Gen(cache_addCmd).FlagCompletion(carapace.ActionMap{
		"package": inLakeWorkspace(lean.ActionLakePackages()),
		"service": lean.ActionLakeCacheServices(),
	})

	carapace.Gen(cache_addCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
