package lean

import (
	"slices"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
)

func installedToolchainNames(output []byte) []string {
	names := make([]string, 0)
	for line := range strings.SplitSeq(string(output), "\n") {
		name := strings.TrimSpace(line)
		if name == "" || name == "no installed toolchains" {
			continue
		}
		names = append(names, strings.Fields(name)[0])
	}
	return names
}

// ActionToolchains completes installed toolchains
//
//	leanprover/lean4:v4.35.0
//	lean4
func ActionToolchains() carapace.Action {
	return carapace.ActionExecCommand("elan", "toolchain", "list")(func(output []byte) carapace.Action {
		return carapace.ActionValues(installedToolchainNames(output)...)
	}).Tag("toolchains").Uid("lean", "toolchain")
}

// ActionReleaseChannels completes release channels
//
//	stable (latest stable release)
//	nightly (latest nightly build)
func ActionReleaseChannels() carapace.Action {
	return carapace.ActionStyledValuesDescribed(
		"beta", "latest release candidate", style.Yellow,
		"nightly", "latest nightly build", style.Red,
		"stable", "latest stable release", style.Green,
	).Tag("release channels")
}

// ActionToolchainSpecifiers completes toolchain specifiers in the form `[<origin>:]<channel>`
// Origins and versions are derived from installed toolchains, no network access is performed.
//
//	stable (latest stable release)
//	leanprover/lean4:v4.35.0
func ActionToolchainSpecifiers() carapace.Action {
	return carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			return carapace.Batch(
				ActionReleaseChannels(),
				actionToolchainOrigins().Suffix(":"),
			).ToA().NoSpace(':')
		case 1:
			return carapace.Batch(
				ActionReleaseChannels(),
				actionInstalledOriginVersions(c.Parts[0]),
			).ToA()
		default:
			return carapace.ActionValues()
		}
	})
}

func actionToolchainOrigins() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		origins := []string{"leanprover/lean4", "leanprover/lean4-nightly"}
		if output, err := c.Command("elan", "toolchain", "list").Output(); err == nil {
			for _, name := range installedToolchainNames(output) {
				if origin, _, found := strings.Cut(name, ":"); found && !slices.Contains(origins, origin) {
					origins = append(origins, origin)
				}
			}
		}
		return carapace.ActionValues(origins...).Tag("origins")
	})
}

func actionInstalledOriginVersions(origin string) carapace.Action {
	return carapace.ActionExecCommand("elan", "toolchain", "list")(func(output []byte) carapace.Action {
		versions := make([]string, 0)
		for _, name := range installedToolchainNames(output) {
			if toolchainOrigin, version, found := strings.Cut(name, ":"); found && toolchainOrigin == origin {
				versions = append(versions, version)
			}
		}
		return carapace.ActionValues(versions...).Tag("installed versions")
	})
}
