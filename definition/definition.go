package definition

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Definition struct {
	Context      string
	Resources    []map[interface{}]interface{}
	ResourceTree *Tree
	Options      map[string]interface{}
}

func New(definition_format interface{}, options map[string]interface{}) *Definition {
	var def *Definition

	switch df_type := definition_format.(type) {
	default:
		log.Fatalf("%s: Is an invalid format type to create a definition", df_type)
	case string:
		def = newFromFile(definition_format.(string), options)
	case map[interface{}]interface{}:
		def = newFromMap(definition_format.(map[interface{}]interface{}), options)
	}

	def.ResourceTree = NewTree(def.Context, def.Resources)

	return def
}

func (d *Definition) Create() {
	tree := d.ResourceTree
	tree.Traverse(func(r Resource) {
		r.Create(d.Options)
	})
}

func (d *Definition) Search(pattern string) Resource {
	resource := make(chan Resource)
	tree := d.ResourceTree

	go func() {
		defer close(resource)
		tree.Traverse(func(r Resource) {
			resource <- r
		})
	}()

	for resource_wanted := range resource {
		pattern := fmt.Sprint(d.Context, "/", pattern)
		if resource_wanted.Id() == pattern {
			return resource_wanted
		}
	}

	return nil
}

func newFromFile(file_name string, options map[string]interface{}) *Definition {
	def := Definition{Options: options}

	definition_content, err := ioutil.ReadFile(file_name)
	if err != nil {
		log.Fatal(err)
	}

	if err = yaml.Unmarshal(definition_content, &def); err != nil {
		log.Fatal(err)
	}

	return &def
}

func newFromMap(map_definition map[interface{}]interface{}, options map[string]interface{}) *Definition {
	return &Definition{
		Context:   map_definition["context"].(string),
		Resources: map_definition["resources"].([]map[interface{}]interface{}),
		Options:   options,
	}
}
