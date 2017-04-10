package cmd

import (
	"fmt"
	"github.com/bronzdoc/symbiote/definition"
	"github.com/bronzdoc/symbiote/vars"
	"github.com/spf13/cobra"
	"log"
)

var createCmd = &cobra.Command{
	Use:   "create TEMPLATE_NAME",
	Short: "Creates a symbiote from a definition",
	Run: func(cmd *cobra.Command, args []string) {
		var template_name string

		if len(args) > 0 {
			template_name = args[0]
		} else {
			err := fmt.Errorf("sym create expects a definition name, see sym create --help")
			log.Fatal(err)
		}

		template_path := fmt.Sprintf("/home/bronzdoc/.symbiote/%s.yml", template_name)

		data, err := vars.Parse(Vars)
		if err != nil {
			log.Fatal(err)
		}

		definition.New(template_path, data).Create()
	},
}

var Vars string

func init() {
	RootCmd.AddCommand(createCmd)
	createCmd.Flags().StringVarP(&Vars, "vars", "", "", "template arguments KEY=VALUE")
}
