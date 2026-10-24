package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/unix/home-manager_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/homemanager"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/nix"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "home-manager",
	Short: "Manage a user environment using Nix",
	Long:  "https://github.com/nix-community/home-manager",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.PersistentFlags().StringS("A", "A", "", "Select an attribute of the configuration file")
	rootCmd.PersistentFlags().StringS("B", "B", "", "Run command to back up existing files rather than fail")
	rootCmd.PersistentFlags().StringArrayS("I", "I", nil, "Add a path to the Nix expression search path")
	rootCmd.PersistentFlags().StringSlice("arg", nil, "Override inputs passed to home-manager.nix")
	rootCmd.PersistentFlags().StringSlice("argstr", nil, "Like --arg but the value is a string")
	rootCmd.PersistentFlags().StringS("b", "b", "", "Move existing files to new path with this extension rather than fail")
	rootCmd.PersistentFlags().String("builders", "", "Remote builders")
	rootCmd.PersistentFlags().Bool("commit-lock-file", false, "Commit changes to the flake's lock file")
	rootCmd.PersistentFlags().String("cores", "", "Threads per job")
	rootCmd.PersistentFlags().Bool("debug", false, "Enable debug output")
	rootCmd.PersistentFlags().BoolP("dry-run", "n", false, "Do a dry run, only prints what actions would be taken")
	rootCmd.PersistentFlags().String("experimental-features", "", "Set the experimental-features setting")
	rootCmd.PersistentFlags().String("extra-experimental-features", "", "Append to the experimental-features setting")
	rootCmd.PersistentFlags().StringP("file", "f", "", "The home configuration file")
	rootCmd.PersistentFlags().String("flake", "", "Use Home Manager configuration at flake-uri")
	rootCmd.PersistentFlags().BoolP("help", "h", false, "Print this help")
	rootCmd.PersistentFlags().Bool("impure", false, "Allow access to mutable paths and repositories")
	rootCmd.PersistentFlags().Bool("keep-failed", false, "Keep temporary directory used by failed builds")
	rootCmd.PersistentFlags().Bool("keep-going", false, "Keep going in case of failed builds")
	rootCmd.PersistentFlags().String("log-format", "", "Set the format of log output")
	rootCmd.PersistentFlags().StringP("max-jobs", "j", "", "Max number of build jobs in parallel")
	rootCmd.PersistentFlags().Bool("no-out-link", false, "Do not create a symlink to the output path")
	rootCmd.PersistentFlags().Bool("no-registries", false, "Don't allow lookups in the flake registries")
	rootCmd.PersistentFlags().Bool("no-substitute", false, "Disable substituters")
	rootCmd.PersistentFlags().Bool("no-update-lock-file", false, "Do not allow any updates to the flake's lock file")
	rootCmd.PersistentFlags().Bool("no-write-lock-file", false, "Do not write the flake's newly generated lock file")
	rootCmd.PersistentFlags().StringSlice("option", nil, "Set the Nix configuration setting name to value")
	rootCmd.PersistentFlags().StringSlice("override-input", nil, "Override a specific flake input")
	rootCmd.PersistentFlags().BoolP("print-build-logs", "L", false, "Print full build logs on standard error")
	rootCmd.PersistentFlags().Bool("recreate-lock-file", false, "Recreate the flake's lock file from scratch")
	rootCmd.PersistentFlags().Bool("refresh", false, "Consider all previously downloaded files out-of-date")
	rootCmd.PersistentFlags().Bool("show-trace", false, "Print stack trace of evaluation errors")
	rootCmd.PersistentFlags().Bool("substitute", false, "Enable substituters")
	rootCmd.PersistentFlags().StringArray("update-input", nil, "Update a specific flake input")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Verbose output")
	rootCmd.PersistentFlags().Bool("version", false, "Print the Home Manager version")

	rootCmd.MarkFlagsMutuallyExclusive("file", "flake")
	rootCmd.MarkFlagsMutuallyExclusive("substitute", "no-substitute")

	rootCmd.Flag("arg").Nargs = 2
	rootCmd.Flag("argstr").Nargs = 2
	rootCmd.Flag("option").Nargs = 2
	rootCmd.Flag("override-input").Nargs = 2

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"A":                           action.ActionConfigurationFileAttributes(rootCmd),
		"B":                           carapace.ActionExecutables(),
		"I":                           carapace.ActionDirectories(),
		"experimental-features":       action.ActionExperimentalFeatures(),
		"extra-experimental-features": action.ActionExperimentalFeatures(),
		"file":                        carapace.ActionFiles(".nix"),
		"flake":                       homemanager.ActionFlakeRefs(),
		"log-format":                  carapace.ActionValues("raw", "internal-json", "bar", "bar-with-logs"),
		"option": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return nix.ActionConfigKeys()
			case 1:
				return nix.ActionConfigValues(c.Parts[0])
			default:
				return carapace.ActionValues()
			}
		}),
		"override-input": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 1:
				return carapace.Batch(
					carapace.ActionDirectories(),
					nix.ActionFlakes(),
				).ToA()
			default:
				return carapace.ActionValues()
			}
		}),
	})
}
