package plan

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

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
	definition_path := fmt.Sprintf("%s/.droid/plans/%s/definition.yml", os.Getenv("HOME"), plan_name)
	definition := definition.New(definition_path, options)
	return New(definition)
}

func (p *Plan) Execute() {
	p.definition.Create()
}

func List() {
	plans_path := fmt.Sprintf("%s/.droid/plans/", os.Getenv("HOME"))
	files, _ := ioutil.ReadDir(plans_path)
	for _, f := range files {
		fmt.Println(f.Name())
	}
}

func Edit(plan_name string) error {
	if !planExists(plan_name) {
		return fmt.Errorf(`plan "%s" doesn't exists.`, plan_name)
	}

	editor := os.Getenv("EDITOR")
	if editor == "" {
		return fmt.Errorf(`$EDITOR is empty, could not edit "%s" plan.`, plan_name)
	}

	definition_path := fmt.Sprintf("%s/.droid/plans/%s/definition.yml", os.Getenv("HOME"), plan_name)

	cmd := exec.Command(editor, definition_path)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		return err
	}

	if err := cmd.Wait(); err != nil {
		return err
	}

	return nil
}

// Get a definition object for a new plan
func PlanDefinition(plan_name string, options map[string]interface{}) *definition.Definition {
	// Default structure for a new plan
	resources := []map[interface{}]interface{}{
		{
			"dir": map[interface{}]interface{}{
				"name": plan_name,
				"dir": map[interface{}]interface{}{
					"name": "templates",
				},
				"files": []interface{}{
					"definition.yml",
				},
			},
		},
	}

	map_definition := map[interface{}]interface{}{
		"context":   fmt.Sprintf("%s/.droid/plans/", os.Getenv("HOME")),
		"resources": resources,
	}

	def := definition.New(map_definition, options)
	resource := def.Search(plan_name + ":" + "templates:definition.yml")

	file := resource.(*definition.File)
	file.SetContent(
		[]byte(`---
context: .
resources:
  - dir:
     name: dir_name_sample
     files:
      - file_name_sample`),
	)

	return def
}

func planExists(plan_name string) bool {
	plans_path := fmt.Sprintf("%s/.droid/plans/", os.Getenv("HOME"))
	files, _ := ioutil.ReadDir(plans_path)
	for _, f := range files {
		if f.Name() == plan_name {
			return true
		}
	}
	return false
}
