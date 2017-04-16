package definition

import (
	"io/ioutil"
	"log"
	"os"
)

type Resource interface {
	Create(map[string]interface{})
	Name() string
	Children() []Resource
	Id() string
}

type Directory struct {
	name     string
	id       string
	children []Resource
}

func (d *Directory) Create(options map[string]interface{}) {
	os.Mkdir(d.id, 0776)
}

func (d *Directory) Name() string {
	return d.name
}

func (d *Directory) Children() []Resource {
	return d.children
}

func (d *Directory) Id() string {
	return d.id
}

type File struct {
	name    string
	id      string
	content []byte
}

func (f *File) Create(options map[string]interface{}) {

	file, err := os.Create(f.id)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// TODO improve how a definition.template and a definition.File interacts
	if f.isContentEmpty() {
		templates_path := options["templates_path"].(string)

		// Check if there is a template for the file
		if f.hasTemplate(templates_path) {
			vars := options["vars"].(map[string]string)
			path := templates_path + "/" + f.name

			content, err := ioutil.ReadFile(path)
			if err != nil {
				log.Fatal(err)
			}

			NewTemplate(f.name, content, vars).Execute(file)
		}
	} else {
		var vars map[string]string
		NewTemplate(f.name, f.content, vars).Execute(file)
	}
}

func (f *File) SetContent(content []byte) {
	f.content = content
}

func (f *File) isContentEmpty() bool {
	return f.content == nil
}

func (f *File) hasTemplate(templates_path string) bool {
	_, err := os.Stat(templates_path + "/" + f.name)
	return err == nil
}

func (f *File) Name() string {
	return f.name
}

func (f *File) Children() []Resource {
	return nil
}

func (f *File) Id() string {
	return f.id
}
