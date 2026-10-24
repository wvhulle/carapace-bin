package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the Lean language server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serveCmd).Standalone()

	rootCmd.AddCommand(serveCmd)

	carapace.Gen(serveCmd).DashAnyCompletion(
		bridge.ActionCarapaceBin("lean"),
	)
}
