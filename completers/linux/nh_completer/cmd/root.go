package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nh",
	Short: "Yet another nix helper",
	Long:  "https://github.com/nix-community/nh",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.PersistentFlags().StringP("elevation-strategy", "e", "", "Choose the privilege elevation strategy")
	rootCmd.PersistentFlags().CountP("quiet", "q", "Decrease logging verbosity")
	rootCmd.PersistentFlags().CountP("verbose", "v", "Increase logging verbosity")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"elevation-strategy": carapace.Batch(
			carapace.ActionValuesDescribed(
				"auto", "Detect doas, sudo, run0 or pkexec",
				"none", "No elevation",
				"passwordless", "Elevate without password prompt on remote hosts",
			).StyleF(style.ForKeyword),
			carapace.ActionFiles(),
		).ToA(),
	})
}
