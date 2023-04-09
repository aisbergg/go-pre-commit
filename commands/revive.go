package commands

import (
	"github.com/spf13/cobra"
)

var revive = &cobra.Command{
	Use:   "revive [files...]",
	Short: "Check your Go source code with revive",
	Long:  ``,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ensureInstalled("revive", "github.com/mgechev/revive")
		runTool("revive", pkgNames(dirNames(args)))
	},
}

func init() {
	rootCmd.AddCommand(revive)
}
