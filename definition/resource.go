package definition

import (
	"log"
	"os"

	"github.com/bronzdoc/symbiote/template"
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
	os.Mkdir(d.id, 0777)
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

	if f.hasTemplate() {
		f.createTemplate(file, options["vars"].(map[string]string))
	}
}

func (f *File) hasTemplate() bool {
	_, err := os.Stat(f.id)
	return err == nil
}

func (f *File) createTemplate(file *os.File, vars map[string]string) {
	_, err := template.New(f.name, vars).Create(file)
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
