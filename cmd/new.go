package cmd

import (
	"fmt"
	"log"

	"github.com/bronzdoc/orbi/plan"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var newCmd = &cobra.Command{
	Use:   "new PLAN",
	Short: "Creates a new plan",
	Run: func(cmd *cobra.Command, args []string) {
		var plan_name string

		if len(args) > 0 {
			plan_name = args[0]
		} else {
			err := fmt.Errorf("orbi plan new expects a plan name, see orbi plan new --help")
			log.Fatal(err)
		}

		// Dinamically plan templates path
		viper.Set("TemplatesPath", fmt.Sprintf(
			"%s/%s/%s", viper.GetString("PlansPath"), viper.GetString("TemplatesDir"), plan_name,
		))

		options := map[string]interface{}{}

		definition := plan.PlanDefinition(plan_name, options)
		plan.New(definition).Execute()
	},
}

func init() {
	planCmd.AddCommand(newCmd)
}
