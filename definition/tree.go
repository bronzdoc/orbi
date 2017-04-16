package definition

type tree struct {
	Root Resource
}

// Creates a definition tree structure
func newTree(context string, definition_resources []map[interface{}]interface{}) *tree {
	resource := &Directory{
		name:     context,
		id:       context,
		children: generate(context, definition_resources),
	}
	return &tree{Root: resource}
}

// Traverse the tree and yield each node to a function
func (t *tree) Traverse(action func(r Resource)) {
	traverse(t.Root, action)
}

// Traverse the tree and yield each resource to a function
func traverse(r Resource, action func(r Resource)) {
	if r.Children() == nil {
		return
	}

	for _, node := range r.Children() {
		action(node)
		traverse(node, action)
	}
}

// Generates a definition Resource hierarchy
func generate(resource_id string, definition_resources []map[interface{}]interface{}) []Resource {
	var resources []Resource
	for _, resource := range definition_resources {
		for key, data := range resource {
			if key == "dir" {
				a_data := data.(map[interface{}]interface{})
				name := a_data["name"].(string)
				id := resource_id + "/" + name

				dir_resources := []map[interface{}]interface{}{
					a_data,
				}

				resources = append(resources, &Directory{
					name:     name,
					id:       id,
					children: generate(id, dir_resources),
				})

			} else if key == "files" {
				a_data := data.([]interface{})
				files := getFileResources(resource_id, filesStringify(a_data))
				resources = append(resources, files...)
			}
		}
	}
	return resources
}

// Convert a []string to []Resource
func getFileResources(resource_id string, file_names []string) []Resource {
	var resources []Resource
	for _, file := range file_names {
		id := resource_id + "/" + file
		resources = append(resources, &File{
			name: file,
			id:   id,
		})
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
