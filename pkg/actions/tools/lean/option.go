package lean

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
)

type leanOptionType int

const (
	leanBoolOption leanOptionType = iota
	leanNatOption
)

type leanOption struct {
	Name        string
	Type        leanOptionType
	Description string
}

var commonLeanOptions = []leanOption{
	{"Elab.async", leanBoolOption, "elaborate theorems and definitions in parallel"},
	{"autoImplicit", leanBoolOption, "automatically bind unbound identifiers as implicit arguments"},
	{"debug.skipKernelTC", leanBoolOption, "skip kernel type checking"},
	{"diagnostics", leanBoolOption, "collect diagnostic information"},
	{"diagnostics.threshold", leanNatOption, "only report diagnostic counters above this threshold"},
	{"experimental.module", leanBoolOption, "enable the module system"},
	{"linter.all", leanBoolOption, "enable all linters"},
	{"linter.constructorNameAsVariable", leanBoolOption, "warn when bound variables share a name with a constructor"},
	{"linter.deprecated", leanBoolOption, "warn on uses of deprecated declarations"},
	{"linter.missingDocs", leanBoolOption, "warn on missing doc strings"},
	{"linter.unnecessarySimpa", leanBoolOption, "warn when simpa can be replaced by simp"},
	{"linter.unusedSimpArgs", leanBoolOption, "warn on unused simp arguments"},
	{"linter.unusedVariables", leanBoolOption, "warn on unused variables"},
	{"maxHeartbeats", leanNatOption, "maximum amount of heartbeats per command (in thousands)"},
	{"maxRecDepth", leanNatOption, "maximum recursion depth"},
	{"maxSynthPendingDepth", leanNatOption, "maximum nesting of pending type class problems"},
	{"pp.all", leanBoolOption, "display coercions, implicit arguments, proof terms, fully qualified names, universes"},
	{"pp.coercions", leanBoolOption, "display coercions as arrows"},
	{"pp.deepTerms", leanBoolOption, "display deeply nested terms"},
	{"pp.explicit", leanBoolOption, "display implicit arguments"},
	{"pp.fieldNotation", leanBoolOption, "use field notation"},
	{"pp.funBinderTypes", leanBoolOption, "display types of lambda parameters"},
	{"pp.maxSteps", leanNatOption, "maximum number of visited expressions before omitting terms"},
	{"pp.mvars", leanBoolOption, "display names of metavariables"},
	{"pp.notation", leanBoolOption, "display notation"},
	{"pp.proofs", leanBoolOption, "display proofs"},
	{"pp.unicode.fun", leanBoolOption, "use `↦` for lambdas"},
	{"pp.universes", leanBoolOption, "display universe levels"},
	{"profiler", leanBoolOption, "show execution times of various Lean components"},
	{"profiler.threshold", leanNatOption, "only report profiling times above this threshold in milliseconds"},
	{"relaxedAutoImplicit", leanBoolOption, "auto-bind any unbound identifier, not only single letters"},
	{"synthInstance.maxHeartbeats", leanNatOption, "maximum amount of heartbeats per type class resolution problem"},
	{"synthInstance.maxSize", leanNatOption, "maximum number of instances used to construct a solution"},
	{"trace.Elab.command", leanBoolOption, "trace command elaboration"},
	{"trace.Elab.step", leanBoolOption, "trace term elaboration steps"},
	{"trace.Meta.Tactic.simp.rewrite", leanBoolOption, "trace simp rewrites"},
	{"trace.Meta.isDefEq", leanBoolOption, "trace definitional equality checks"},
	{"trace.Meta.synthInstance", leanBoolOption, "trace type class resolution"},
	{"trace.profiler", leanBoolOption, "activate nested traces with execution time above threshold"},
	{"trace.profiler.threshold", leanNatOption, "threshold in milliseconds for trace.profiler"},
	{"warningAsError", leanBoolOption, "treat warnings as errors"},
}

// ActionOptions completes common configuration option names (see `set_option`)
//
//	autoImplicit (automatically bind unbound identifiers as implicit arguments)
//	maxHeartbeats (maximum amount of heartbeats per command (in thousands))
func ActionOptions() carapace.Action {
	values := make([]string, 0, len(commonLeanOptions)*2)
	for _, option := range commonLeanOptions {
		values = append(values, option.Name, option.Description)
	}
	return carapace.ActionValuesDescribed(values...).Tag("options")
}

// ActionOptionValues completes values of the given configuration option
//
//	true
//	false
func ActionOptionValues(name string) carapace.Action {
	for _, option := range commonLeanOptions {
		if option.Name == name && option.Type == leanBoolOption {
			return carapace.ActionStyledValues(
				"true", style.Green,
				"false", style.Red,
			)
		}
	}
	return carapace.ActionValues()
}

// ActionOptionAssignments completes configuration option assignments in the form `name=value`
//
//	autoImplicit=false
//	maxHeartbeats=400000
func ActionOptionAssignments() carapace.Action {
	return carapace.ActionMultiParts("=", func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			return ActionOptions().Suffix("=")
		case 1:
			return ActionOptionValues(c.Parts[0])
		default:
			return carapace.ActionValues()
		}
	})
}
