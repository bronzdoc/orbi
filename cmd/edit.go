package cmd

import (
	"fmt"
	"log"

	"github.com/bronzdoc/droid/plan"
	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit PLAN_NAME",
	Short: "Edit a plan definition, $EDITOR will be used to open the file.",
	Run: func(cmd *cobra.Command, args []string) {
		var plan_name string

		if len(args) > 0 {
			plan_name = args[0]
		} else {
			err := fmt.Errorf("droid plan edit expects a plan name, see droid plan edit --help")
			log.Fatal(err)
		}

		if err := plan.Edit(plan_name); err != nil {
			log.Fatalf("edit command failed: %v\n", err)
		}
	},
}

func init() {
	planCmd.AddCommand(editCmd)
}
