package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/bronzdoc/droid/plan"
	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new PLAN",
	Short: "Creates a new plan",
	Run: func(cmd *cobra.Command, args []string) {
		var plan_name string

		if len(args) > 0 {
			plan_name = args[0]
		} else {
			err := fmt.Errorf("droid plan new expects a plan name, see droid plan new --help")
			log.Fatal(err)
		}

		options := map[string]interface{}{
			"templates_path": fmt.Sprintf(
				"%s/.droid/plans/%s/templates", os.Getenv("HOME"), plan_name,
			),
		}

		definition := plan.PlanDefinition(plan_name, options)
		plan.New(definition).Execute()
	},
}

func init() {
	planCmd.AddCommand(newCmd)
}
