package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/lean"
	"github.com/spf13/cobra"
)

var cache_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Download build outputs from a remote service into the Lake cache",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cache_getCmd).Standalone()

	cache_getCmd.Flags().Bool("download-arts", false, "download both mappings and artifacts")
	cache_getCmd.Flags().Bool("force-download", false, "redownload existing files")
	cache_getCmd.Flags().Bool("mappings-only", false, "only download mappings, delay artifacts")
	cache_getCmd.Flags().String("max-revs", "", "backtrack up to n revisions")
	cache_getCmd.Flags().String("package", "", "fetch outputs for set package")
	cache_getCmd.Flags().String("platform", "", "with Reservoir or --repo, set the platform")
	cache_getCmd.Flags().String("repo", "", "GitHub repository of the package or its fork")
	cache_getCmd.Flags().String("rev", "", "lookup artifacts only for set revision")
	cache_getCmd.Flags().String("scope", "", "verbatim scope for a custom endpoint")
	cache_getCmd.Flags().String("service", "", "fetch outputs from set cache service")
	cache_getCmd.Flags().String("toolchain", "", "with Reservoir or --repo, set the toolchain")
	cacheCmd.AddCommand(cache_getCmd)

	cache_getCmd.Flag("download-arts").Hidden = true

	carapace.Gen(cache_getCmd).FlagCompletion(carapace.ActionMap{
		"package": inLakeWorkspace(lean.ActionLakePackages()),
		"service": lean.ActionLakeCacheServices(),
	})

	carapace.Gen(cache_getCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
