package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/traverse"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "lake",
	Short: "Lean build system and package manager",
	Long:  "https://lean-lang.org/doc/reference/latest/Build-Tools-and-Distribution/Lake/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func inLakeWorkspace(action carapace.Action) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return action.ChdirF(traverse.Flag(rootCmd.Flag("dir")))
	})
}

func actionLakeLogLevels() carapace.Action {
	return carapace.ActionValues("trace", "info", "warning", "error")
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.PersistentFlags().StringArrayS("K", "K", nil, "set the configuration file option named key")
	rootCmd.PersistentFlags().BoolS("U", "U", false, "deprecated shorthand for --update")
	rootCmd.PersistentFlags().Bool("allow-empty", false, "accept bare builds with no default targets configured")
	rootCmd.PersistentFlags().Bool("ansi", false, "use ANSI escape codes to prettify output")
	rootCmd.PersistentFlags().StringP("dir", "d", "", "use the package configuration in a specific directory")
	rootCmd.PersistentFlags().Bool("fail-fast", false, "stop scheduling new build jobs after the first required failure")
	rootCmd.PersistentFlags().String("fail-level", "", "minimum log level to fail a build")
	rootCmd.PersistentFlags().StringP("file", "f", "", "use a specific file for the package configuration")
	rootCmd.PersistentFlags().BoolP("help", "h", false, "print help of the program or a command and exit")
	rootCmd.PersistentFlags().Bool("iofail", false, "fail build if any I/O or other info is logged")
	rootCmd.PersistentFlags().BoolP("json", "J", false, "output JSON-formatted results (in `lake query`)")
	rootCmd.PersistentFlags().Bool("keep-toolchain", false, "do not update toolchain on workspace update")
	rootCmd.PersistentFlags().String("log-level", "", "minimum log level to output on success")
	rootCmd.PersistentFlags().Bool("no-ansi", false, "do not use ANSI escape codes to prettify output")
	rootCmd.PersistentFlags().Bool("no-build", false, "exit immediately if a build target is not up-to-date")
	rootCmd.PersistentFlags().Bool("no-cache", false, "build packages locally; do not download build caches")
	rootCmd.PersistentFlags().Bool("offline", false, "do not access the network")
	rootCmd.PersistentFlags().Bool("old", false, "only rebuild modified modules (ignore transitive deps)")
	rootCmd.PersistentFlags().String("packages", "", "JSON file of package entries that override the manifest")
	rootCmd.PersistentFlags().BoolP("quiet", "q", false, "hide informational logs and the progress indicator")
	rootCmd.PersistentFlags().BoolP("reconfigure", "R", false, "elaborate configuration files instead of using OLeans")
	rootCmd.PersistentFlags().BoolP("rehash", "H", false, "hash all files for traces (do not trust `.hash` files)")
	rootCmd.PersistentFlags().Bool("text", false, "output results as plain text (in `lake query`)")
	rootCmd.PersistentFlags().Bool("try-cache", false, "attempt to download build caches for supported packages")
	rootCmd.PersistentFlags().Bool("update", false, "update dependencies on load (e.g., before a build)")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "show trace logs (command invocations) and built targets")
	rootCmd.Flags().Bool("version", false, "print version and exit")
	rootCmd.PersistentFlags().Bool("wfail", false, "fail build if warnings are logged")

	rootCmd.Flag("U").Hidden = true
	rootCmd.Flag("offline").Hidden = true

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"dir":        carapace.ActionDirectories(),
		"fail-level": actionLakeLogLevels(),
		"file":       carapace.ActionFiles(".lean", ".toml"),
		"log-level":  actionLakeLogLevels(),
		"packages":   carapace.ActionFiles(".json"),
	})
}
