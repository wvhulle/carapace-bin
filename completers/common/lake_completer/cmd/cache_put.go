package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/lean"
	"github.com/spf13/cobra"
)

var cache_putCmd = &cobra.Command{
	Use:   "put",
	Short: "Upload build outputs from the Lake cache to a remote service",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cache_putCmd).Standalone()

	cache_putCmd.Flags().String("package", "", "upload for set package")
	cache_putCmd.Flags().String("platform", "", "with --repo, sets the platform")
	cache_putCmd.Flags().String("repo", "", "scope w/ repository + toolchain & platform")
	cache_putCmd.Flags().String("scope", "", "upload under set scope verbatim")
	cache_putCmd.Flags().String("service", "", "upload to set cache service")
	cache_putCmd.Flags().String("toolchain", "", "with --repo, sets the toolchain")
	cacheCmd.AddCommand(cache_putCmd)

	carapace.Gen(cache_putCmd).FlagCompletion(carapace.ActionMap{
		"package": inLakeWorkspace(lean.ActionLakePackages()),
		"service": lean.ActionLakeCacheServices(),
	})

	carapace.Gen(cache_putCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
