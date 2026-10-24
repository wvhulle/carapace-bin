package homemanager

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
)

type homeGeneration struct {
	id        string
	createdAt string
	storePath string
	current   bool
}

var generationLinePattern = regexp.MustCompile(`^(\S+ \S+) : id (\S+) -> (\S+)( \(current\))?$`)

func parseGenerations(output []byte) []homeGeneration {
	generations := make([]homeGeneration, 0)
	for _, line := range strings.Split(string(output), "\n") {
		if match := generationLinePattern.FindStringSubmatch(line); match != nil {
			generations = append(generations, homeGeneration{
				id:        match[2],
				createdAt: match[1],
				storePath: match[3],
				current:   match[4] != "",
			})
		}
	}
	return generations
}

// ActionGenerations completes home environment generations
// The current generation is highlighted.
//
//	2629 (2026-06-20 11:51)
//	2630 (2026-06-21 09:12)
func ActionGenerations() carapace.Action {
	return carapace.ActionExecCommand("home-manager", "generations")(func(output []byte) carapace.Action {
		vals := make([]string, 0)
		for _, generation := range parseGenerations(output) {
			generationStyle := style.Default
			if generation.current {
				generationStyle = style.Carapace.KeywordPositive
			}
			vals = append(vals, generation.id, generation.createdAt+" "+generation.storePath, generationStyle)
		}
		return carapace.ActionStyledValuesDescribed(vals...)
	}).Tag("generations")
}

// ActionSpecialisations completes specialisations of the current home environment generation
//
//	dark
//	light
func ActionSpecialisations() carapace.Action {
	return carapace.ActionExecCommand("home-manager", "generations")(func(output []byte) carapace.Action {
		for _, generation := range parseGenerations(output) {
			if !generation.current {
				continue
			}
			entries, err := os.ReadDir(filepath.Join(generation.storePath, "specialisation"))
			if err != nil {
				return carapace.ActionValues()
			}
			names := make([]string, 0, len(entries))
			for _, entry := range entries {
				names = append(names, entry.Name())
			}
			return carapace.ActionValues(names...)
		}
		return carapace.ActionValues()
	}).Tag("specialisations")
}
