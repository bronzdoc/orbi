package definition

import (
	"github.com/bronzdoc/droid/template"
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
	name string
	id   string
}

func (f *File) Create(options map[string]interface{}) {
	file, err := os.Create(f.id)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Check if there is a template for the file
	templates_path := options["templates_path"].(string)
	if f.hasTemplate(templates_path) {
		vars := options["vars"].(map[string]string)
		f.createTemplate(
			file,
			templates_path,
			vars,
		)
	}
}

func (f *File) hasTemplate(templates_path string) bool {
	_, err := os.Stat(templates_path + "/" + f.name)
	return err == nil
}

func (f *File) createTemplate(file *os.File, templates_path string, vars map[string]string) {
	_, err := template.New(f.name, templates_path, vars).Create(file)
	if err != nil {
		log.Fatal(err)
	}
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
