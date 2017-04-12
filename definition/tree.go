package definition

import r "github.com/bronzdoc/symbiote/definition/resource"

type tree struct {
	root r.Resource
}

func NewTree(kv_resources []map[interface{}]interface{}) *tree {
	resource := &r.Directory{
		Name:     ".",
		Children: generate(kv_resources),
	}
	return &tree{root: resource}
}

// Generates a r.Resource hierarchy
func generate(kv_resources []map[interface{}]interface{}) []r.Resource {
	var resources []r.Resource
	for _, resource := range kv_resources {
		for key, data := range resource {
			if key == "dir" {
				a_data := data.(map[interface{}]interface{})
				name := a_data["name"].(string)

				dir_resources := []map[interface{}]interface{}{
					a_data,
				}

				resources = append(resources, &r.Directory{
					Name:     name,
					Children: generate(dir_resources),
				})

			} else if key == "files" {
				a_data := data.([]interface{})
				files := getFileResources(filesStringify(a_data))
				resources = append(resources, files...)
			}
		}
	}
	return resources
}

// Convert a []string to []r.Resource
func getFileResources(file_names []string) []r.Resource {
	var resources []r.Resource
	for _, file := range file_names {
		resources = append(resources, &r.File{Name: file})
	}
	return resources
}

// Convert a []interface{} to []string
func filesStringify(file_names []interface{}) []string {
	var files []string
	for _, file_name := range file_names {

		files = append(files, file_name.(string))
	}
	return files
}
