package definition

import (
	"github.com/bronzdoc/symbiote/template"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

type definition struct {
	Context   string
	Resources []map[string]interface{}
	vars      map[string]string
}

func New(definition_file string, vars map[string]string) *definition {
	def := definition{}
	def.vars = vars
	data, err := ioutil.ReadFile(definition_file)

	if err != nil {
		log.Fatal(err)
	}

	if err = yaml.Unmarshal(data, &def); err != nil {
		log.Fatal(err)
	}

	return &def
}

func (s *definition) Create() {
	resources := s.Resources
	context := s.Context

	for i := range resources {
		for key, val := range resources[i] {
			if key == "dir" {
				data := val.(map[interface{}]interface{})
				dir_name := data["name"]
				os.Mkdir(context+"/"+dir_name.(string), 0777)
				if _, ok := data["files"]; ok {
					files := data["files"].([]interface{})
					for j := range files {
						filename := files[j].(string)
						file_path := context + "/" + dir_name.(string) + "/" + filename
						f := create_file(file_path)
						defer f.Close()
						template_path := os.Getenv("HOME") + "/" + ".definitions/templates" + "/" + filename
						_, err := template.New(filename, template_path, s.vars).Create(f)
						if err != nil {
							log.Fatal(err)
						}
					}
				}
			}
		}
	}
}

func create_file(filename string) *os.File {
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	return f
}
