package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cache_servicesCmd = &cobra.Command{
	Use:   "services",
	Short: "Print configured remote cache services",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cache_servicesCmd).Standalone()

	cacheCmd.AddCommand(cache_servicesCmd)
}
