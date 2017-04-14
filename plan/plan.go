package plan

import (
	"fmt"
	"github.com/bronzdoc/droid/definition"
)

type Plan struct {
	definition *definition.Definition
}

func New(definition *definition.Definition) *Plan {
	return &Plan{
		definition: definition,
	}
}

func PlanFactory(plan_name string, options map[string]interface{}) *Plan {
	// TODO this should be in a config object
	definition_path := fmt.Sprintf("/home/bronzdoc/.droid/plans/%s/definition.yml", plan_name)
	definition := definition.New(definition_path, options)
	return New(definition)
}

func (p *Plan) Execute() {
	p.definition.Create()
}
