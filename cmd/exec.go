package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/bronzdoc/droid/plan"
	"github.com/bronzdoc/droid/vars"
	"github.com/spf13/cobra"
)

var execCmd = &cobra.Command{
	Use:   "exec PLAN",
	Short: "Executes a plan",
	Run: func(cmd *cobra.Command, args []string) {
		var plan_name string

		if len(args) > 0 {
			plan_name = args[0]
		} else {
			err := fmt.Errorf("droid exec expects a plan name, see droid exec --help")
			log.Fatal(err)
		}

		vars, err := vars.Parse(Vars)
		if err != nil {
			log.Fatal(err)
		}

		options := map[string]interface{}{
			"vars": vars,
			"templates_path": fmt.Sprintf(
				"%s/.droid/plans/%s/templates", os.Getenv("HOME"), plan_name,
			),
		}

		plan := plan.PlanFactory(plan_name, options)
		plan.Execute()
	},
}

var Vars string

func init() {
	RootCmd.AddCommand(execCmd)
	execCmd.Flags().StringVarP(&Vars, "vars", "", "", "template vars KEY=VALUE")
}
