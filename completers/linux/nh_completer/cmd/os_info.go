package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var os_infoCmd = &cobra.Command{
	Use:   "info",
	Short: "List available generations from profile path",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(os_infoCmd).Standalone()

	os_infoCmd.Flags().String("fields", "", "Comma-delimited list of field(s) to display")
	os_infoCmd.Flags().StringP("profile", "P", "/nix/var/nix/profiles/system", "Path to Nix profiles directory")

	carapace.Gen(os_infoCmd).FlagCompletion(carapace.ActionMap{
		"fields": carapace.ActionValuesDescribed(
			"confRev", "Configuration Revision",
			"date", "Build Date",
			"id", "Generation Id",
			"kernel", "Kernel Version",
			"nver", "NixOS Version",
			"size", "Closure Size",
			"spec", "Specialisations",
		).UniqueList(","),
		"profile": carapace.ActionFiles(),
	})

	osCmd.AddCommand(os_infoCmd)
}
