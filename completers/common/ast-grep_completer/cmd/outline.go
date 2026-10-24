package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/ast-grep_completer/cmd/action"
	"github.com/spf13/cobra"
)

var outlineCmd = &cobra.Command{
	Use:   "outline",
	Short: "Explore code structure for symbols, imports, exports, and members",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(outlineCmd).Standalone()

	outlineCmd.Flags().String("color", "", "Controls output color")
	outlineCmd.Flags().Bool("follow", false, "Follow symbolic links")
	outlineCmd.Flags().StringSlice("globs", nil, "Include or exclude file paths")
	outlineCmd.Flags().BoolP("help", "h", false, "Print help")
	outlineCmd.Flags().String("items", "", "Select which top-level items to include")
	outlineCmd.Flags().String("json", "", "Output outline entries in structured JSON")
	outlineCmd.Flags().StringP("lang", "l", "", "Specify the input language")
	outlineCmd.Flags().String("match", "", "Keep only top-level items matching this regex")
	outlineCmd.Flags().Bool("no-default-outline-rules", false, "Do not load bundled outline extractor definitions")
	outlineCmd.Flags().StringSlice("no-ignore", nil, "Do not respect hidden file system or ignore files")
	outlineCmd.Flags().String("outline-rules", "", "Load additional outline extractor definitions")
	outlineCmd.Flags().Bool("pub-members", false, "Display only public members in member views")
	outlineCmd.Flags().Bool("stdin", false, "Enable search code from StdIn")
	outlineCmd.Flags().StringP("threads", "j", "", "Set the approximate number of threads to use")
	outlineCmd.Flags().String("type", "", "Keep only top-level items with these comma-separated symbol types")
	outlineCmd.Flags().String("view", "", "Select the text presentation")

	outlineCmd.Flag("json").NoOptDefVal = " "
	rootCmd.AddCommand(outlineCmd)

	carapace.Gen(outlineCmd).FlagCompletion(carapace.ActionMap{
		"color": action.ActionColorModes(),
		"items": carapace.ActionValuesDescribed(
			"auto", "Use `structure` for file or stdin input, `exports` when any directory is given",
			"structure", "Top-level items defined locally in the file, excluding imports",
			"exports", "Top-level items exported from the file or module",
			"imports", "Top-level items imported from other files or modules",
			"all", "All top-level items, including imports and exports",
		),
		"json":          action.ActionJsonStyles(),
		"lang":          action.ActionLanguages(),
		"no-ignore":     action.ActionIgnoreFileTypes(),
		"outline-rules": carapace.ActionFiles(),
		"view": carapace.ActionValuesDescribed(
			"auto", "Use `digest` for file or stdin input, `names` when any directory is given",
			"names", "One grouped name line per symbol type for each file",
			"signatures", "One source/signature line per top-level item",
			"digest", "Signatures plus compact member name digests",
			"expanded", "Signatures plus one source/signature line per direct member",
		),
	})

	carapace.Gen(outlineCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
