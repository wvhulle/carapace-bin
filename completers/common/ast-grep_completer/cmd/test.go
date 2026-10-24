package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/ast-grep_completer/cmd/action"
	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Test ast-grep rules",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(testCmd).Standalone()

	testCmd.Flags().String("color", "", "Controls output color")
	testCmd.Flags().StringP("filter", "f", "", "Only run rule test cases that matches REGEX")
	testCmd.Flags().Bool("follow", false, "Follow symbolic links while searching test YAML files")
	testCmd.Flags().BoolP("help", "h", false, "Print help")
	testCmd.Flags().Bool("include-off", false, "Include `severity:off` rules in test")
	testCmd.Flags().BoolP("interactive", "i", false, "Start an interactive review to update snapshots selectively")
	testCmd.Flags().Bool("skip-snapshot-tests", false, "Only check if the test code is valid, without checking rule output")
	testCmd.Flags().String("snapshot-dir", "", "Specify the directory name storing snapshots")
	testCmd.Flags().StringP("test-dir", "t", "", "the directories to search test YAML files")
	testCmd.Flags().BoolP("update-all", "U", false, "Update the content of all snapshots that have changed in test")
	rootCmd.AddCommand(testCmd)

	carapace.Gen(testCmd).FlagCompletion(carapace.ActionMap{
		"color":        action.ActionColorModes(),
		"snapshot-dir": carapace.ActionDirectories(),
		"test-dir":     carapace.ActionDirectories(),
	})
}
