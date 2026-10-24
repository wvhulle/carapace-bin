package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/lean"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "lean",
	Short: "Lean theorem prover and programming language",
	Long:  "https://lean-lang.org/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringArrayS("D", "D", nil, "set a configuration option (see set_option command)")
	rootCmd.Flags().StringP("bc", "b", "", "name of the LLVM bitcode file")
	rootCmd.Flags().StringP("c", "c", "", "name of the C output file")
	rootCmd.Flags().Bool("deps", false, "just print dependencies of a Lean input")
	rootCmd.Flags().Bool("deps-json", false, "print dependencies of Lean inputs as JSON")
	rootCmd.Flags().StringP("error", "E", "", "report Lean messages of kind as errors")
	rootCmd.Flags().Bool("exitOnPanic", false, "exit on panic")
	rootCmd.Flags().Bool("features", false, "display features compiler provides (eg. LLVM support)")
	rootCmd.Flags().BoolP("githash", "g", false, "display the git commit hash number used to build this binary")
	rootCmd.Flags().BoolP("help", "h", false, "display this message")
	rootCmd.Flags().StringP("i", "i", "", "create ilean file")
	rootCmd.Flags().String("incr-header-save", "", "EXPERIMENTAL: like `--incr-save`, but save only the header (state after importing)")
	rootCmd.Flags().String("incr-load", "", "EXPERIMENTAL: reuse a snapshot saved by `--incr-(header-)save` at start of run")
	rootCmd.Flags().String("incr-save", "", "EXPERIMENTAL: save a full incremental snapshot of post-elaboration state at end of run")
	rootCmd.Flags().Bool("json", false, "report Lean output (e.g., messages) as JSON (one per line)")
	rootCmd.Flags().String("load-dynlib", "", "load shared library to make its symbols available to the interpreter")
	rootCmd.Flags().StringP("memory", "M", "", "maximum amount of memory that should be used by Lean (in megabytes)")
	rootCmd.Flags().StringP("o", "o", "", "create olean file")
	rootCmd.Flags().String("plugin", "", "load and initialize Lean shared library for registering linters etc.")
	rootCmd.Flags().Bool("print-libdir", false, "print the installation directory for Lean's built-in libraries and exit")
	rootCmd.Flags().Bool("print-prefix", false, "print the installation prefix for Lean and exit")
	rootCmd.Flags().Bool("profile", false, "display elaboration/type checking time for each definition/theorem")
	rootCmd.Flags().BoolP("quiet", "q", false, "do not print verbose messages")
	rootCmd.Flags().StringP("root", "R", "", "set package root directory from which the module name of the input file is calculated")
	rootCmd.Flags().Bool("run", false, "call the 'main' definition in the given file with the remaining arguments")
	rootCmd.Flags().Bool("server", false, "start lean in server mode")
	rootCmd.Flags().String("setup", "", "JSON file with module setup data (supersedes the file's header)")
	rootCmd.Flags().BoolP("short-version", "V", false, "display short version number")
	rootCmd.Flags().Bool("src-deps", false, "just print dependency sources of a Lean input")
	rootCmd.Flags().Bool("stats", false, "display environment statistics")
	rootCmd.Flags().Bool("stdin", false, "take input from stdin")
	rootCmd.Flags().StringP("threads", "j", "", "number of threads used to process lean files")
	rootCmd.Flags().StringP("timeout", "T", "", "maximum number of memory allocations per task")
	rootCmd.Flags().StringP("trust", "t", "", "trust level (default: max) 0 means do not trust any macro")
	rootCmd.Flags().StringP("tstack", "s", "", "thread stack size in Kb")
	rootCmd.Flags().BoolP("version", "v", false, "display version information")
	rootCmd.Flags().Bool("worker", false, "start lean in server-worker mode")

	rootCmd.Flag("deps-json").Hidden = true
	rootCmd.Flag("exitOnPanic").Hidden = true

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"D":                lean.ActionOptionAssignments(),
		"bc":               carapace.ActionFiles(".bc"),
		"c":                carapace.ActionFiles(".c"),
		"i":                carapace.ActionFiles(".ilean"),
		"incr-header-save": carapace.ActionFiles(),
		"incr-load":        carapace.ActionFiles(),
		"incr-save":        carapace.ActionFiles(),
		"load-dynlib":      carapace.ActionFiles(".so", ".dylib", ".dll"),
		"o":                carapace.ActionFiles(".olean"),
		"plugin":           carapace.ActionFiles(".so", ".dylib", ".dll"),
		"root":             carapace.ActionDirectories(),
		"setup":            carapace.ActionFiles(".json"),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(".lean"),
	)

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if rootCmd.Flag("run").Changed {
				return carapace.ActionFiles()
			}
			return carapace.ActionFiles(".lean")
		}),
	)
}
