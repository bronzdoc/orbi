package definition

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type definition struct {
	Context      string
	Resources    []map[interface{}]interface{}
	ResourceTree *tree
	Options      map[string]interface{}
}

func New(definition_file string, options map[string]interface{}) *definition {
	def := definition{Options: options}

	definition_content, err := ioutil.ReadFile(definition_file)
	if err != nil {
		log.Fatal(err)
	}

	if err = yaml.Unmarshal(definition_content, &def); err != nil {
		log.Fatal(err)
	}

	def.ResourceTree = newTree(def.Context, def.Resources)

	return &def
}

func (d *definition) Create() {
	tree := d.ResourceTree
	tree.Traverse(func(node Resource) {
		node.Create(d.Options)
	})
}
