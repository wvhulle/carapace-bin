package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var envCmd = &cobra.Command{
	Use:   "env",
	Short: "Execute a command in Lake's environment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(envCmd).Standalone()
	envCmd.Flags().SetInterspersed(false)

	rootCmd.AddCommand(envCmd)

	carapace.Gen(envCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin(),
	)
}
