package cmd

import (
	"fmt"
	"github.com/bronzdoc/symbiote/definition"
	"github.com/bronzdoc/symbiote/vars"
	"github.com/spf13/cobra"
	"log"
)

var createCmd = &cobra.Command{
	Use:   "create DEFINITION_NAME",
	Short: "Creates a symbiote from a definition",
	Run: func(cmd *cobra.Command, args []string) {
		var definition_name string

		if len(args) > 0 {
			definition_name = args[0]
		} else {
			err := fmt.Errorf("sym create expects a definition name, see sym create --help")
			log.Fatal(err)
		}

		// TODO this should be in a config object
		definition_path := fmt.Sprintf("/home/bronzdoc/.droid/%s.yml", definition_name)

		vars, err := vars.Parse(Vars)
		if err != nil {
			log.Fatal(err)
		}

		options := map[string]interface{}{"vars": vars}
		definition.New(definition_path, options).Create()
	},
}

var Vars string

func init() {
	RootCmd.AddCommand(createCmd)
	createCmd.Flags().StringVarP(&Vars, "vars", "", "", "definition arguments KEY=VALUE")
}
