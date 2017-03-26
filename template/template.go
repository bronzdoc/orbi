package template

import (
	"io"
	"io/ioutil"
	"os"
	"text/template"
)

type Template struct {
	name string
	path string
	vars map[string]interface{}
}

func New(name, template_path string) *Template {
	return &Template{
		name: name,
		path: template_path,
	}
}

func (t *Template) Create(w io.Writer, vars map[string]string) (*Template, error) {
	if _, err := os.Stat(t.path); err == nil {
		tmpl_content, err := ioutil.ReadFile(t.path)
		if err != nil {
			return t, err
		}

		tmpl, err := template.New(t.name).Option("missingkey=error").Parse(string(tmpl_content))
		if err != nil {
			return t, err
		}

		err = tmpl.Execute(w, vars)
		if err != nil {
			return t, err
		}
	}

	return t, nil
}
