package cmd

import (
	"fmt"
	"log"

	"github.com/bronzdoc/orbi/plan"
	"github.com/bronzdoc/orbi/vars"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var execCmd = &cobra.Command{
	Use:   "exec PLAN",
	Short: "Executes a plan",
	Run: func(cmd *cobra.Command, args []string) {
		var plan_name string

		if len(args) > 0 {
			plan_name = args[0]
		} else {
			err := fmt.Errorf("orbi exec expects a plan name, see orbi exec --help")
			log.Fatal(err)
		}

		vars, err := vars.Parse(Vars)
		if err != nil {
			log.Fatal(err)
		}

		// Dinamically plan templates path
		viper.Set("TemplatesPath", fmt.Sprintf(
			"%s/%s/%s", viper.GetString("PlansPath"), viper.GetString("TemplateDir"), plan_name,
		))

		options := map[string]interface{}{
			"vars": vars,
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
