package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/ast-grep_completer/cmd/action"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ast-grep",
	Short: "Search and Rewrite code at large scale using AST pattern",
	Long:  "https://ast-grep.github.io/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.PersistentFlags().StringP("config", "c", "", "Path to ast-grep root config, default is sgconfig.yml")
	rootCmd.PersistentFlags().String("custom-languages", "", "Control how sgconfig.yml handles native custom language libraries")
	rootCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.Flags().BoolP("version", "V", false, "Print version")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"config":           carapace.ActionFiles(".yml", ".yaml"),
		"custom-languages": action.ActionCustomLanguagePolicies(),
	})
}
