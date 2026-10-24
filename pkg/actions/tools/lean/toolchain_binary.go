package lean

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionElanProxiedBinaries completes the binaries elan proxies
//
//	lake
//	lean
func ActionElanProxiedBinaries() carapace.Action {
	return carapace.ActionValues("lake", "lean", "leanc", "leanchecker", "leanmake", "leanpkg").Tag("binaries")
}

// ActionToolchainBinaries completes the binaries of the active toolchain
// Falls back to the binaries elan proxies when the active toolchain can not be resolved.
//
//	lake
//	leanchecker
func ActionToolchainBinaries() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		output, err := c.Command("elan", "which", "lean").Output()
		if err != nil {
			return ActionElanProxiedBinaries()
		}
		binDir := filepath.Dir(strings.TrimSpace(string(output)))
		entries, err := os.ReadDir(binDir)
		if err != nil {
			return ActionElanProxiedBinaries()
		}
		binaries := make([]string, 0, len(entries))
		for _, entry := range entries {
			if !entry.IsDir() {
				binaries = append(binaries, entry.Name())
			}
		}
		return carapace.ActionValues(binaries...).Tag("binaries")
	})
}
