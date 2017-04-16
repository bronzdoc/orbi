package definition

import (
	"io/ioutil"
	"log"
	"strings"

	"gopkg.in/yaml.v2"
)

type Definition struct {
	Context      string
	Resources    []map[interface{}]interface{}
	ResourceTree *tree
	Options      map[string]interface{}
}

func New(definition_format interface{}, options map[string]interface{}) *Definition {
	var def *Definition

	switch df_type := definition_format.(type) {
	default:
		log.Fatalf("%s: Is an invalid format type to create a definition", df_type)
	case string:
		def = newFromFileName(definition_format.(string), options)
	case map[interface{}]interface{}:
		def = newFromMap(definition_format.(map[interface{}]interface{}), options)
	}

	def.ResourceTree = newTree(def.Context, def.Resources)

	return def
}

func (d *Definition) Create() {
	tree := d.ResourceTree
	tree.Traverse(func(r Resource) {
		r.Create(d.Options)
	})
}

func (d *Definition) Search(pattern string) Resource {
	strings_to_match := strings.Split(pattern, ":")
	match_counter := 0
	var resource_wanted Resource

	tree := d.ResourceTree
	tree.Traverse(func(r Resource) {
		if r.Name() == strings_to_match[match_counter] {
			match_counter += 1
		} else {
			match_counter = 0
		}

		if match_counter == len(strings_to_match) {
			resource_wanted = r
			return
		}
	})
	return resource_wanted
}

func newFromFileName(file_name string, options map[string]interface{}) *Definition {
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
