package actions

import (
	_ "fmt"
	"github.com/bronzdoc/skeletor/template"
	"io/ioutil"
	"log"
	"os"
)

func Create(template template.Template) {
	skeleton := template.Skeleton
	context := template.Context

	for i := range skeleton {
		for key, val := range skeleton[i] {
			if key == "dir" {
				data := val.(map[interface{}]interface{})
				dir_name := data["name"]
				os.Mkdir(context+"/"+dir_name.(string), 0777)
				if _, ok := data["files"]; ok {
					files := data["files"].([]interface{})
					for j := range files {
						filename := context + "/" + dir_name.(string) + "/" + files[j].(string)
						f := create_file(filename)
						defer f.Close()
						template_name := os.Getenv("HOME") + "/" + ".skeletor/templates" + "/" + files[j].(string)
						if _, err := os.Stat(template_name); err == nil {
							add_template_content_to_file(template_name, f)
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

func add_template_content_to_file(template_name string, file *os.File) {
	template_content, err := ioutil.ReadFile(template_name)
	if err != nil {
		log.Fatal(err)
	}
	file.Write(template_content)
}
