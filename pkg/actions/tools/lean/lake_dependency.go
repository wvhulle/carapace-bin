package lean

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
)

// ActionLakeDependencies completes the dependencies locked in `lake-manifest.json`
//
//	mathlib (master)
//	batteries (main)
func ActionLakeDependencies() carapace.Action {
	return actionWithLakeWorkspace(func(workspace lakeWorkspace) carapace.Action {
		values := make([]string, 0)
		for _, dependency := range workspace.Manifest.Packages {
			description := dependency.URL
			if dependency.InputRev != nil {
				description = *dependency.InputRev
			} else if dependency.Type == "path" {
				description = dependency.Dir
			}
			values = append(values, dependency.Name, description)
		}
		return carapace.ActionValuesDescribed(values...)
	}).Tag("dependencies").Uid("lean", "lake-dependency")
}

// ActionLakePackages completes the root package and the dependencies of the workspace
//
//	compfiles (root package)
//	mathlib (master)
func ActionLakePackages() carapace.Action {
	return actionWithLakeWorkspace(func(workspace lakeWorkspace) carapace.Action {
		batch := carapace.Batch(ActionLakeDependencies())
		if workspace.Root.Name != "" {
			batch = append(batch, carapace.ActionStyledValuesDescribed(workspace.Root.Name, "root package", style.Bold))
		}
		return batch.ToA()
	}).Tag("packages")
}
