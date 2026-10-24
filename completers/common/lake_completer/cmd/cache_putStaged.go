package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/lean"
	"github.com/spf13/cobra"
)

var cache_putStagedCmd = &cobra.Command{
	Use:   "put-staged",
	Short: "Upload build outputs from a staging directory to a remote service",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cache_putStagedCmd).Standalone()

	cache_putStagedCmd.Flags().String("platform", "", "with --repo, set the platform")
	cache_putStagedCmd.Flags().String("repo", "", "scope w/ repository + toolchain & platform")
	cache_putStagedCmd.Flags().String("rev", "", "upload for set revision")
	cache_putStagedCmd.Flags().String("scope", "", "upload under set scope verbatim")
	cache_putStagedCmd.Flags().String("service", "", "upload to set cache service")
	cache_putStagedCmd.Flags().String("toolchain", "", "with --repo, set the toolchain")
	cacheCmd.AddCommand(cache_putStagedCmd)

	carapace.Gen(cache_putStagedCmd).FlagCompletion(carapace.ActionMap{
		"service": lean.ActionLakeCacheServices(),
	})

	carapace.Gen(cache_putStagedCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
