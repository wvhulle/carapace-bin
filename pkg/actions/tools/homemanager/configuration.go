package homemanager

import (
	"encoding/json"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/nix"
)

func resolveLocalFlake(c carapace.Context, flake string) string {
	if !strings.HasPrefix(flake, ".") && !strings.HasPrefix(flake, "/") && !strings.HasPrefix(flake, "~") {
		return flake
	}
	if absoluteFlake, err := c.Abs(flake); err == nil {
		return absoluteFlake
	}
	return flake
}

// ActionConfigurations completes homeConfigurations of given flake
//
//	alice@laptop
//	bob
func ActionConfigurations(flake string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommandE("nix", "eval", "--json", resolveLocalFlake(c, flake)+"#homeConfigurations", "--apply", "builtins.attrNames")(func(output []byte, err error) carapace.Action {
			if err != nil {
				return carapace.ActionMessage("failed to evaluate homeConfigurations of %v", flake)
			}
			var configurationNames []string
			if err := json.Unmarshal(output, &configurationNames); err != nil {
				return carapace.ActionMessage(err.Error())
			}
			return carapace.ActionValues(configurationNames...)
		})
	}).Tag("home configurations")
}

// ActionFlakeRefs completes a flake reference with a homeConfigurations attribute
//
//	.#alice@laptop
//	github:alice/dotfiles#bob
func ActionFlakeRefs() carapace.Action {
	return carapace.ActionMultiPartsN("#", 2, func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			return carapace.Batch(
				carapace.ActionDirectories(),
				nix.ActionFlakes().Suffix("#"),
			).ToA()
		default:
			return ActionConfigurations(c.Parts[0])
		}
	})
}
