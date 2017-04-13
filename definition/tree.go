package definition

type tree struct {
	root Resource
}

func NewTree(kv_resources []map[interface{}]interface{}) *tree {
	resource := &Directory{
		name:     ".",
		children: generate(kv_resources),
	}
	return &tree{root: resource}
}

func (t *tree) Traverse(action func(r Resource)) {
	traverse(t.root, action)
}

// Traverse the tree and yield each node to a function
func traverse(r Resource, action func(r Resource)) {
	if r.Children() == nil {
		return
	}

	for _, node := range r.Children() {
		action(node)
		traverse(node, action)
	}
}

// Generates a Resource hierarchy
func generate(kv_resources []map[interface{}]interface{}) []Resource {
	var resources []Resource
	for _, resource := range kv_resources {
		for key, data := range resource {
			if key == "dir" {
				a_data := data.(map[interface{}]interface{})
				name := a_data["name"].(string)

				dir_resources := []map[interface{}]interface{}{
					a_data,
				}

				resources = append(resources, &Directory{
					name:     name,
					children: generate(dir_resources),
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

// Convert a []string to []Resource
func getFileResources(file_names []string) []Resource {
	var resources []Resource
	for _, file := range file_names {
		resources = append(resources, &File{name: file})
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
