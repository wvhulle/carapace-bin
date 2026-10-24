package action

import (
	"os"
	"path/filepath"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/nix"
	"github.com/spf13/cobra"
)

func findDefaultConfigurationFile(c carapace.Context) (string, bool) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", false
	}
	configHome := c.Getenv("XDG_CONFIG_HOME")
	if configHome == "" {
		configHome = filepath.Join(homeDir, ".config")
	}
	for _, candidate := range []string{
		filepath.Join(configHome, "home-manager", "home.nix"),
		filepath.Join(configHome, "nixpkgs", "home.nix"),
		filepath.Join(homeDir, ".nixpkgs", "home.nix"),
	} {
		if _, err := os.Stat(candidate); err == nil {
			return candidate, true
		}
	}
	return "", false
}

func ActionConfigurationFileAttributes(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		configurationFile := cmd.Flag("file").Value.String()
		if configurationFile == "" {
			defaultFile, found := findDefaultConfigurationFile(c)
			if !found {
				return carapace.ActionValues()
			}
			configurationFile = defaultFile
		}
		includePaths, _ := cmd.Flags().GetStringArray("I")
		include := ""
		if len(includePaths) > 0 {
			include = includePaths[0]
		}
		return nix.ActionAttributes(nix.AttributeOpts{
			Source:  configurationFile,
			Include: include,
		})
	})
}

func ActionExperimentalFeatures() carapace.Action {
	return carapace.ActionValues(
		"auto-allocate-uids",
		"ca-derivations",
		"cgroups",
		"dynamic-derivations",
		"fetch-closure",
		"fetch-tree",
		"flakes",
		"git-hashing",
		"impure-derivations",
		"local-overlay-store",
		"mounted-ssh-store",
		"nix-command",
		"no-url-literals",
		"pipe-operators",
		"read-only-local-store",
		"recursive-nix",
		"verified-fetches",
	).UniqueList(" ")
}
