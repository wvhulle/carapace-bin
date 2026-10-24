package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var shakeCmd = &cobra.Command{
	Use:   "shake",
	Short: "Minimize imports in Lean source files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(shakeCmd).Standalone()

	shakeCmd.Flags().Bool("add-public", false, "Add new imports as `public` if they were in the original public closure")
	shakeCmd.Flags().Bool("explain", false, "Show which constants require each import")
	shakeCmd.Flags().Bool("fix", false, "Apply suggested fixes directly to source files")
	shakeCmd.Flags().Bool("force", false, "Skip the `lake build --no-build` sanity check")
	shakeCmd.Flags().Bool("gh-style", false, "Output in GitHub problem matcher format")
	shakeCmd.Flags().Bool("keep-implied", false, "Preserve imports implied by other imports")
	shakeCmd.Flags().Bool("keep-prefix", false, "Prefer parent module imports over specific submodules")
	shakeCmd.Flags().Bool("keep-public", false, "Preserve all `public` imports for API stability")
	shakeCmd.Flags().StringArray("only", nil, "minimize only this module")
	shakeCmd.Flags().Bool("trace", false, "trace the analysis")
	rootCmd.AddCommand(shakeCmd)

	shakeCmd.Flag("only").Hidden = true
	shakeCmd.Flag("trace").Hidden = true
}
