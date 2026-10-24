package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var newsCmd = &cobra.Command{
	Use:   "news",
	Short: "Show news entries in a pager",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(newsCmd).Standalone()

	rootCmd.AddCommand(newsCmd)
}
