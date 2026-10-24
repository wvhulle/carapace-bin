package lean

import (
	"github.com/carapace-sh/carapace"
)

// ActionLakeScripts completes scripts declared in `lakefile.lean` of the workspace and its dependencies
// Scripts of dependencies are prefixed with their package name.
//
//	update (script)
//	mathlib/test (script of mathlib)
func ActionLakeScripts() carapace.Action {
	return actionWithLakeWorkspace(func(workspace lakeWorkspace) carapace.Action {
		values := make([]string, 0)
		for _, script := range workspace.Root.targetsOfKind(lakeScriptKind) {
			values = append(values, script, "script")
		}
		for packageName, config := range workspace.dependencyConfigs() {
			for _, script := range config.targetsOfKind(lakeScriptKind) {
				values = append(values, packageName+"/"+script, "script of "+packageName)
			}
		}
		return carapace.ActionValuesDescribed(values...)
	}).Tag("scripts").Uid("lean", "lake-script")
}
