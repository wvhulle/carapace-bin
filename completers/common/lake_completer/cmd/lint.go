package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lintCmd = &cobra.Command{
	Use:   "lint",
	Short: "Lint the workspace's root package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lintCmd).Standalone()

	lintCmd.Flags().Bool("builtin-lint", false, "run builtin environment and text linters")
	lintCmd.Flags().Bool("builtin-only", false, "run only builtin linters, skip the lint driver")
	lintCmd.Flags().String("checks", "", "comma-separated list of workspace modules providing package code quality checks")
	lintCmd.Flags().Bool("code-quality", false, "record each linter warning as a code quality check result")
	lintCmd.Flags().String("lint-only", "", "like `--linters`, but report only the linters the spec positively enables")
	lintCmd.Flags().String("linters", "", "override linter options for the lint build")
	lintCmd.Flags().Bool("record-exceptions", false, "record each linter warning as a `set_option <linter> false in` exception")
	rootCmd.AddCommand(lintCmd)

	carapace.Gen(lintCmd).DashAnyCompletion(
		carapace.ActionFiles(),
	)
}
