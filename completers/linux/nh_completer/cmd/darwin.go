package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/nh_completer/cmd/common"
	"github.com/spf13/cobra"
)

var darwinCmd = &cobra.Command{
	Use:   "darwin",
	Short: "Nix-darwin functionality",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(darwinCmd).Standalone()

	rootCmd.AddCommand(darwinCmd)
}

func addDarwinRebuildFlags(cmd *cobra.Command) {
	common.AddCommonRebuildFlags(cmd)
	common.AddUpdateFlags(cmd)
	common.AddRemoteBuildHostFlag(cmd)

	cmd.Flags().BoolP("bypass-root-check", "R", false, "Don't panic if calling nh as root")
	cmd.Flags().StringP("hostname", "H", "", "Select this hostname from darwinConfigurations")
	cmd.Flags().Bool("show-activation-logs", false, "Show activation logs")
}
