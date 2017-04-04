package skeleton

import (
	"github.com/bronzdoc/skeletor/template"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

type skeleton struct {
	Context   string
	Resources []map[string]interface{}
	vars      map[string]string
}

func New(skeleton_file string, vars map[string]string) *skeleton {
	sk := skeleton{}
	sk.vars = vars
	data, err := ioutil.ReadFile(skeleton_file)

	if err != nil {
		log.Fatal(err)
	}

	if err = yaml.Unmarshal(data, &sk); err != nil {
		log.Fatal(err)
	}

	return &sk
}

func (s *skeleton) Create() {
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
						template_path := os.Getenv("HOME") + "/" + ".skeletor/templates" + "/" + filename
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
