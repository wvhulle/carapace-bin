package action

import "github.com/carapace-sh/carapace"

func ActionLanguages() carapace.Action {
	return carapace.ActionValues(
		"bash",
		"c",
		"cpp",
		"csharp",
		"css",
		"dart",
		"elixir",
		"go",
		"haskell",
		"hcl",
		"html",
		"java",
		"javascript",
		"json",
		"kotlin",
		"lua",
		"markdown",
		"nix",
		"php",
		"python",
		"ruby",
		"rust",
		"scala",
		"solidity",
		"swift",
		"tsx",
		"typescript",
		"yaml",
		"zig",
	)
}

func ActionJsonStyles() carapace.Action {
	return carapace.ActionValuesDescribed(
		"pretty", "Prints the matches as a pretty-printed JSON array",
		"stream", "Prints each match as a separate JSON object per line",
		"compact", "Prints the matches as a single-line JSON array",
	)
}

func ActionColorModes() carapace.Action {
	return carapace.ActionValuesDescribed(
		"auto", "Try to use colors, but don't force the issue",
		"always", "Try very hard to emit colors",
		"ansi", "Like always, but only emit ANSI color codes",
		"never", "Never emit colors",
	)
}

func ActionIgnoreFileTypes() carapace.Action {
	return carapace.ActionValuesDescribed(
		"hidden", "Search hidden files and directories",
		"dot", "Don't respect .ignore files",
		"exclude", "Don't respect manually configured repository ignore files",
		"global", "Don't respect ignore files from global sources",
		"parent", "Don't respect ignore files in parent directories",
		"vcs", "Don't respect version control ignore files",
	)
}

func ActionInspectGranularities() carapace.Action {
	return carapace.ActionValuesDescribed(
		"nothing", "Do not show any tracing information",
		"summary", "Show summary about how many files are scanned and skipped",
		"entity", "Show per-file/per-rule tracing information",
	)
}

func ActionCustomLanguagePolicies() carapace.Action {
	return carapace.ActionValuesDescribed(
		"allow", "Load native custom language libraries for this invocation",
		"ignore", "Skip native custom language libraries without reporting an error",
		"trust", "Load native custom language libraries and remember this project",
		"revoke", "Forget this project and ignore its native custom language libraries",
	)
}
