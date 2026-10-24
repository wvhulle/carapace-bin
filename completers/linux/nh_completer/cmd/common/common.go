package common

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/nix"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

func ActionDiffModes() carapace.Action {
	return carapace.ActionValuesDescribed(
		"always", "Always display package diff",
		"auto", "Display package diff only if the current and the deployed configuration match",
		"never", "Never display package diff",
	).StyleF(style.ForKeyword)
}

func ActionInstallable(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if cmd.Flag("file").Changed || cmd.Flag("expr").Changed {
			return carapace.ActionValues()
		}
		return carapace.Batch(
			carapace.ActionDirectories(),
			nix.ActionFlakes(),
		).ToA()
	})
}

func AddInstallableFlags(cmd *cobra.Command) {
	cmd.Flags().StringP("expr", "E", "", "Nix expression with an optional attribute path")
	cmd.Flags().StringP("file", "f", "", "Path to file with an optional attribute path")

	cmd.MarkFlagsMutuallyExclusive("expr", "file")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"file": carapace.ActionFiles(),
	})

	carapace.Gen(cmd).PositionalCompletion(
		ActionInstallable(cmd),
	)
}

func AddCommonRebuildFlags(cmd *cobra.Command) {
	cmd.Flags().BoolP("ask", "a", false, "Ask for confirmation")
	cmd.Flags().StringP("diff", "d", "auto", "Whether to display a package diff")
	cmd.Flags().BoolP("dry", "n", false, "Only print actions, without performing them")
	cmd.Flags().Bool("no-nom", false, "Don't use nix-output-monitor for the build process")
	cmd.Flags().StringP("out-link", "o", "", "Path to save the result link")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"diff":     ActionDiffModes(),
		"out-link": carapace.ActionFiles(),
	})

	AddNixBuildPassthroughFlags(cmd)
	AddInstallableFlags(cmd)
}

func AddNixBuildPassthroughFlags(cmd *cobra.Command) {
	cmd.Flags().Bool("accept-flake-config", false, "Accept configuration from flakes")
	cmd.Flags().String("builders", "", "Explicitly define remote builders")
	cmd.Flags().Bool("commit-lock-file", false, "Commit the lock file after updates")
	cmd.Flags().String("cores", "", "Number of cores Nix should utilize")
	cmd.Flags().Bool("fallback", false, "Attempt to build locally if substituters fail")
	cmd.Flags().Bool("impure", false, "Allow impure builds")
	cmd.Flags().StringArrayP("include", "I", nil, "Paths to include")
	cmd.Flags().Bool("json", false, "Output results in JSON format")
	cmd.Flags().BoolP("keep-failed", "K", false, "Keep build outputs from failed builds")
	cmd.Flags().BoolP("keep-going", "k", false, "Continue building despite encountering errors")
	cmd.Flags().String("log-format", "", "Logging format used by Nix")
	cmd.Flags().StringP("max-jobs", "j", "", "Number of concurrent jobs Nix should run")
	cmd.Flags().BoolP("no-build-output", "Q", false, "Suppress build output")
	cmd.Flags().Bool("no-net", false, "Prohibit network usage")
	cmd.Flags().Bool("no-registries", false, "Do not use registries (deprecated)")
	cmd.Flags().Bool("no-update-lock-file", false, "Do not update the flake.lock file")
	cmd.Flags().Bool("no-use-registries", false, "Do not use registries")
	cmd.Flags().Bool("no-write-lock-file", false, "Do not write a lock file")
	cmd.Flags().Bool("offline", false, "Build without internet access")
	cmd.Flags().StringArray("option", nil, "Set a Nix configuration option")
	cmd.Flags().StringArray("override-input", nil, "Override a specific flake input")
	cmd.Flags().BoolP("print-build-logs", "L", false, "Print build logs directly to stdout")
	cmd.Flags().Bool("recreate-lock-file", false, "Recreate the flake.lock file entirely")
	cmd.Flags().Bool("refresh", false, "Refresh flakes to the latest revision")
	cmd.Flags().Bool("repair", false, "Repair corrupted store paths")
	cmd.Flags().BoolP("show-trace", "t", false, "Display tracebacks on errors")
	cmd.Flags().Bool("use-substitutes", false, "Use substitutes when copying")

	cmd.Flag("option").Nargs = 2
	cmd.Flag("override-input").Nargs = 2

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"include":    carapace.ActionDirectories(),
		"log-format": carapace.ActionValues("bar", "bar-with-logs", "internal-json", "raw"),
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

func AddUpdateFlags(cmd *cobra.Command) {
	cmd.Flags().BoolP("update", "u", false, "Update all flake inputs")
	cmd.Flags().StringArrayP("update-input", "U", nil, "Update the specified flake input(s)")

	cmd.MarkFlagsMutuallyExclusive("update", "update-input")
}

func AddRemoteBuildHostFlag(cmd *cobra.Command) {
	cmd.Flags().String("build-host", "", "Build the configuration on a different host over SSH")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"build-host": net.ActionHosts(),
	})
}

func AddOsRebuildFlags(cmd *cobra.Command) {
	AddCommonRebuildFlags(cmd)
	AddUpdateFlags(cmd)
	AddRemoteBuildHostFlag(cmd)

	cmd.Flags().BoolP("bypass-root-check", "R", false, "Don't panic if calling nh as root")
	cmd.Flags().StringP("hostname", "H", "", "Select this hostname from nixosConfigurations")
	cmd.Flags().Bool("install-bootloader", false, "Install bootloader for switch and boot commands")
	cmd.Flags().BoolP("no-specialisation", "S", false, "Ignore specialisations")
	cmd.Flags().Bool("no-validate", false, "Skip pre-activation system validation checks")
	cmd.Flags().StringP("specialisation", "s", "", "Explicitly select some specialisation")
	cmd.Flags().String("target-host", "", "Deploy the built configuration to a different host over SSH")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"target-host": net.ActionHosts(),
	})
}
