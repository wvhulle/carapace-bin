package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/ast-grep_completer/cmd/action"
	"github.com/spf13/cobra"
)

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan and rewrite code by configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(scanCmd).Standalone()

	scanCmd.Flags().StringP("after", "A", "", "Show NUM lines after each match")
	scanCmd.Flags().StringP("before", "B", "", "Show NUM lines before each match")
	scanCmd.Flags().String("color", "", "Controls output color")
	scanCmd.Flags().StringP("context", "C", "", "Show NUM lines around each match")
	scanCmd.Flags().StringSlice("error", nil, "Set rule severity to error")
	scanCmd.Flags().Bool("files-with-matches", false, "Print only the paths with at least one match and suppress match contents")
	scanCmd.Flags().String("filter", "", "Scan the codebase with rules with ids matching REGEX")
	scanCmd.Flags().Bool("follow", false, "Follow symbolic links")
	scanCmd.Flags().String("format", "", "Output warning/error messages in different formats")
	scanCmd.Flags().StringSlice("globs", nil, "Include or exclude file paths")
	scanCmd.Flags().BoolP("help", "h", false, "Print help")
	scanCmd.Flags().StringSlice("hint", nil, "Set rule severity to hint")
	scanCmd.Flags().Bool("include-metadata", false, "Include rule metadata in the json output")
	scanCmd.Flags().StringSlice("info", nil, "Set rule severity to info")
	scanCmd.Flags().String("inline-rules", "", "Scan the codebase with a rule defined by the provided RULE_TEXT")
	scanCmd.Flags().String("inspect", "", "Inspect information for file/rule discovery and scanning")
	scanCmd.Flags().BoolP("interactive", "i", false, "Start interactive edit session")
	scanCmd.Flags().String("json", "", "Output matches in structured JSON")
	scanCmd.Flags().String("max-results", "", "Show at most NUM results and stop running once the limit is reached")
	scanCmd.Flags().String("min-severity", "", "Set the minimum severity of rules to scan")
	scanCmd.Flags().StringSlice("no-ignore", nil, "Do not respect hidden file system or ignore files")
	scanCmd.Flags().StringSlice("off", nil, "Turn off rule")
	scanCmd.Flags().String("report-style", "", "Diagnostic report style")
	scanCmd.Flags().StringP("rule", "r", "", "Scan the codebase with the single rule located at the path RULE_FILE")
	scanCmd.Flags().Bool("stdin", false, "Enable search code from StdIn")
	scanCmd.Flags().StringP("threads", "j", "", "Set the approximate number of threads to use")
	scanCmd.Flags().BoolP("update-all", "U", false, "Apply all rewrite without confirmation if true")
	scanCmd.Flags().StringSlice("warning", nil, "Set rule severity to warning")

	scanCmd.Flag("error").NoOptDefVal = " "
	scanCmd.Flag("hint").NoOptDefVal = " "
	scanCmd.Flag("info").NoOptDefVal = " "
	scanCmd.Flag("json").NoOptDefVal = " "
	scanCmd.Flag("off").NoOptDefVal = " "
	scanCmd.Flag("warning").NoOptDefVal = " "
	rootCmd.AddCommand(scanCmd)

	carapace.Gen(scanCmd).FlagCompletion(carapace.ActionMap{
		"color":        action.ActionColorModes(),
		"format":       carapace.ActionValues("github", "sarif"),
		"inspect":      action.ActionInspectGranularities(),
		"json":         action.ActionJsonStyles(),
		"min-severity": carapace.ActionValues("error", "warning", "info", "hint", "off"),
		"no-ignore":    action.ActionIgnoreFileTypes(),
		"report-style": carapace.ActionValuesDescribed(
			"rich", "Output a richly formatted diagnostic, with source code previews",
			"medium", "Output a condensed diagnostic, with a line number, severity, message and notes",
			"short", "Output a short diagnostic, with a line number, severity, and message",
		),
		"rule": carapace.ActionFiles(".yml", ".yaml"),
	})

	carapace.Gen(scanCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
