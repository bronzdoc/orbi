package cmd

import (
	"fmt"
	"log"

	"github.com/bronzdoc/orbi/plan"
	"github.com/bronzdoc/orbi/vars"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var varsFlag string

var execCmd = &cobra.Command{
	Use:   "exec PLAN",
	Short: "Executes a plan",
	Run: func(cmd *cobra.Command, args []string) {
		var planName string

		if len(args) > 0 {
			planName = args[0]
		} else {
			err := fmt.Errorf("orbi exec expects a plan name, see orbi exec --help")
			log.Fatal(err)
		}

		vars, err := vars.Parse(varsFlag)
		if err != nil {
			log.Fatal(err)
		}

		// Dinamically plan templates path
		viper.Set("TemplatesPath", fmt.Sprintf(
			"%s/%s/%s", viper.GetString("PlansPath"), planName, viper.GetString("TemplatesDir"),
		))

		options := map[string]interface{}{
			"vars": vars,
		}

		plan := plan.Factory(planName, options)

		if err := plan.Execute(); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(execCmd)
	execCmd.Flags().StringVarP(&varsFlag, "vars", "", "", "template vars KEY=VALUE")
}
