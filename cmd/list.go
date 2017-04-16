package cmd

import (
	"github.com/bronzdoc/droid/plan"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List plans",
	Run: func(cmd *cobra.Command, args []string) {
		plan.List()
	},
}

func init() {
	planCmd.AddCommand(listCmd)
}
