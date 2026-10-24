package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var search_packagesCmd = &cobra.Command{
	Use:   "packages",
	Short: "Search packages via search.nixos.org",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(search_packagesCmd).Standalone()

	addSearchBackendFlags(search_packagesCmd)
	search_packagesCmd.Flags().BoolP("json", "j", false, "Output results as JSON")
	search_packagesCmd.Flags().BoolP("platforms", "P", false, "Show supported platforms for each package")

	searchCmd.AddCommand(search_packagesCmd)
}
