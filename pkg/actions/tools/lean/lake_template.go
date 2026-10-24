package lean

import (
	"github.com/carapace-sh/carapace"
)

// ActionLakeTemplates completes package templates of `lake new` and `lake init`
// Templates can be suffixed with the configuration language.
//
//	std (library and executable; default)
//	math.lean
func ActionLakeTemplates() carapace.Action {
	return carapace.ActionMultiParts(".", func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			return carapace.ActionValuesDescribed(
				"exe", "executable only",
				"lib", "library only",
				"math", "library with Mathlib standards for linting and workflows",
				"math-lax", "library only with a Mathlib dependency",
				"std", "library and executable; default",
			)
		case 1:
			return ActionLakeConfigLanguages()
		default:
			return carapace.ActionValues()
		}
	}).Tag("templates")
}

// ActionLakeConfigLanguages completes languages of the package configuration
//
//	lean (lakefile.lean)
//	toml (lakefile.toml)
func ActionLakeConfigLanguages() carapace.Action {
	return carapace.ActionValuesDescribed(
		"lean", "lakefile.lean",
		"toml", "lakefile.toml",
	).Tag("configuration languages")
}
