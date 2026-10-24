package lean

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
)

var lakeBuildableKinds = []lakeTargetKind{
	lakeLibraryKind,
	lakeExecutableKind,
	lakeExternLibKind,
	lakeCustomKind,
	lakeInputFileKind,
	lakeInputDirKind,
}

func actionLakeConfigTargets(config lakePackageConfig, kinds ...lakeTargetKind) carapace.Action {
	values := make([]string, 0)
	for _, target := range config.Targets {
		for _, kind := range kinds {
			if target.Kind == kind {
				values = append(values, target.Name, string(target.Kind))
			}
		}
	}
	return carapace.ActionValuesDescribed(values...)
}

// ActionLakeExecutables completes executables of the workspace and its dependencies
//
//	cache (executable of mathlib)
//	checkSolution (executable)
func ActionLakeExecutables() carapace.Action {
	return actionWithLakeWorkspace(func(workspace lakeWorkspace) carapace.Action {
		values := make([]string, 0)
		for _, executable := range workspace.Root.targetsOfKind(lakeExecutableKind) {
			values = append(values, executable, "executable")
		}
		for packageName, config := range workspace.dependencyConfigs() {
			for _, executable := range config.targetsOfKind(lakeExecutableKind) {
				values = append(values, executable, "executable of "+packageName)
			}
		}
		return carapace.ActionValuesDescribed(values...)
	}).Tag("executables").Uid("lean", "lake-executable")
}

// ActionLakeFacets completes library and module facets of build targets
//
//	leanArts (Lean artifacts (*.olean, *.ilean, *.c files))
//	olean (OLean (binary blob of Lean data for importers))
func ActionLakeFacets() carapace.Action {
	return carapace.ActionValuesDescribed(
		"bc", "compiled LLVM bitcode file",
		"bc.o", "compiled object file (of its LLVM bitcode file)",
		"c", "compiled C file",
		"c.o", "compiled object file (of its C file)",
		"deps", "dependencies (e.g., imports, shared libraries, etc.)",
		"dynlib", "shared library (e.g., for `--load-dynlib`)",
		"ilean", "ILean (binary blob of metadata for the Lean LSP server)",
		"leanArts", "Lean artifacts (*.olean, *.ilean, *.c files)",
		"o", "compiled object file (of its configured backend)",
		"olean", "OLean (binary blob of Lean data for importers)",
		"shared", "shared artifact (*.so, *.dll, or *.dylib file)",
		"static", "static artifact (*.a file)",
	).Tag("facets")
}

// ActionLakeTargets completes build targets in the form `[@[<package>]/][<target>|<file>][:<facet>]`
// Targets are read from the package configurations of the workspace and its dependencies.
//
//	Compfiles (library)
//	@mathlib/Mathlib (library)
func ActionLakeTargets() carapace.Action {
	return carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			if strings.HasPrefix(c.Value, "@") {
				return actionLakePackageScopedTargets()
			}
			return carapace.Batch(
				actionWithLakeWorkspace(func(workspace lakeWorkspace) carapace.Action {
					return actionLakeConfigTargets(workspace.Root, lakeBuildableKinds...)
				}),
				carapace.ActionValues("@").NoSpace('@'),
				carapace.ActionFiles(".lean"),
			).ToA()
		case 1:
			return ActionLakeFacets()
		default:
			return carapace.ActionValues()
		}
	}).Tag("targets").Uid("lean", "lake-target")
}

func actionLakePackageScopedTargets() carapace.Action {
	return actionWithLakeWorkspace(func(workspace lakeWorkspace) carapace.Action {
		return carapace.ActionMultiParts("/", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				packageNames := []string{"@"}
				if workspace.Root.Name != "" {
					packageNames = append(packageNames, "@"+workspace.Root.Name)
				}
				for _, dependency := range workspace.Manifest.Packages {
					packageNames = append(packageNames, "@"+dependency.Name)
				}
				return carapace.ActionValues(packageNames...).Style(style.Blue).Tag("packages")
			case 1:
				packageName := strings.TrimPrefix(c.Parts[0], "@")
				if packageName == "" || packageName == workspace.Root.Name {
					return actionLakeConfigTargets(workspace.Root, lakeBuildableKinds...)
				}
				if config, ok := workspace.dependencyConfigs()[packageName]; ok {
					return actionLakeConfigTargets(config, lakeBuildableKinds...)
				}
				return carapace.ActionValues()
			default:
				return carapace.ActionValues()
			}
		})
	})
}
