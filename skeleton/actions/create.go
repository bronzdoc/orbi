package actions

import (
	"fmt"
	"github.com/bronzdoc/skeletor/template"
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
				fmt.Println(dir_name)
				os.Mkdir(context+"/"+dir_name.(string), 0777)
				if _, ok := data["files"]; ok {
					files := data["files"].([]interface{})
					for j := range files {
						filename := context + "/" + dir_name.(string) + "/" + files[j].(string)
						fmt.Println(filename)
						f, err := os.Create(filename)
						if err != nil {
							log.Fatal(err)
						}
						defer f.Close()
					}
				}
			}
		}
	}
}
