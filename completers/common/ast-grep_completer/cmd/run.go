package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/ast-grep_completer/cmd/action"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run one time search or rewrite in command line",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runCmd).Standalone()

	runCmd.Flags().StringP("after", "A", "", "Show NUM lines after each match")
	runCmd.Flags().StringP("before", "B", "", "Show NUM lines before each match")
	runCmd.Flags().String("color", "", "Controls output color")
	runCmd.Flags().StringP("context", "C", "", "Show NUM lines around each match")
	runCmd.Flags().String("debug-query", "", "Print query pattern's tree-sitter AST")
	runCmd.Flags().Bool("files-with-matches", false, "Print only the paths with at least one match and suppress match contents")
	runCmd.Flags().Bool("follow", false, "Follow symbolic links")
	runCmd.Flags().StringSlice("globs", nil, "Include or exclude file paths")
	runCmd.Flags().String("heading", "", "Controls whether to print the file name as heading")
	runCmd.Flags().BoolP("help", "h", false, "Print help")
	runCmd.Flags().String("inspect", "", "Inspect information for file/rule discovery and scanning")
	runCmd.Flags().BoolP("interactive", "i", false, "Start interactive edit session")
	runCmd.Flags().String("json", "", "Output matches in structured JSON")
	runCmd.Flags().StringP("kind", "k", "", "AST kind to match")
	runCmd.Flags().StringP("lang", "l", "", "The language of the pattern")
	runCmd.Flags().StringSlice("no-ignore", nil, "Do not respect hidden file system or ignore files")
	runCmd.Flags().StringP("pattern", "p", "", "AST pattern to match")
	runCmd.Flags().StringP("rewrite", "r", "", "String to replace the matched AST node")
	runCmd.Flags().String("selector", "", "AST kind to extract sub-part of pattern to match")
	runCmd.Flags().Bool("stdin", false, "Enable search code from StdIn")
	runCmd.Flags().String("strictness", "", "The strictness of the pattern")
	runCmd.Flags().StringP("threads", "j", "", "Set the approximate number of threads to use")
	runCmd.Flags().BoolP("update-all", "U", false, "Apply all rewrite without confirmation if true")

	runCmd.Flag("debug-query").NoOptDefVal = " "
	runCmd.Flag("json").NoOptDefVal = " "
	rootCmd.AddCommand(runCmd)

	carapace.Gen(runCmd).FlagCompletion(carapace.ActionMap{
		"color": action.ActionColorModes(),
		"debug-query": carapace.ActionValuesDescribed(
			"pattern", "Print the query parsed in Pattern format",
			"ast", "Print the query in tree-sitter AST format, only named nodes are shown",
			"cst", "Print the query in tree-sitter CST format, both named and unnamed nodes are shown",
			"sexp", "Print the query in S-expression format",
		),
		"heading": carapace.ActionValuesDescribed(
			"auto", "Print heading for terminal tty but not for piped output",
			"always", "Always print heading regardless of output type",
			"never", "Never print heading regardless of output type",
		),
		"inspect":   action.ActionInspectGranularities(),
		"json":      action.ActionJsonStyles(),
		"lang":      action.ActionLanguages(),
		"no-ignore": action.ActionIgnoreFileTypes(),
		"strictness": carapace.ActionValuesDescribed(
			"cst", "Match all nodes exactly",
			"smart", "Match all nodes except source trivial nodes",
			"ast", "Match only ast nodes",
			"relaxed", "Match ast nodes except comments",
			"signature", "Match ast nodes except comments, without text",
			"template", "Similar to smart but match text only, node kinds are ignored",
		),
	})

	carapace.Gen(runCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
