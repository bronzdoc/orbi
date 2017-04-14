package template

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"text/template"
)

type Template struct {
	name string
	path string
	vars map[string]string
}

func New(name, template_path string, vars map[string]string) *Template {
	return &Template{
		name: name,
		path: template_path,
		vars: vars,
	}
}

func (t *Template) Create(w io.Writer) (*Template, error) {
	if _, err := os.Stat(t.path); err == nil {
		tmpl_content, err := ioutil.ReadFile(t.path)
		if err != nil {
			return t, err
		}

		var_names, err := findVars(string(tmpl_content))
		if err != nil {
			return t, err
		}

		var missing_vars []string
		for _, name := range var_names {
			key := name
			if _, ok := t.vars[key]; !ok {
				missing_vars = append(missing_vars, key)
			}
		}

		if len(missing_vars) > 0 {
			var err_msg bytes.Buffer
			for _, missing_var := range missing_vars {
				err_msg.WriteString("{{" + "." + missing_var + "}}\n")
			}

			return t, fmt.Errorf(
				"%s: Missing template variables:\n%s",
				t.path,
				err_msg.String(),
			)
		}

		tmpl, err := template.New(t.name).Option("missingkey=error").Parse(
			string(tmpl_content),
		)
		if err != nil {
			return t, err
		}

		err = tmpl.Execute(w, t.vars)
		if err != nil {
			return t, err
		}
	}

	return t, nil
}

func findVars(template_content string) ([]string, error) {
	var vars []string

	regex, err := regexp.Compile(`\s?{{\.[a-zA-Z0-9][a-zA-Z0-9]*}}`)
	if err != nil {
		return vars, err
	}

	matches := regex.FindAllString(template_content, -1)

	regex, err = regexp.Compile(`[a-zA-Z0-0]+`)
	if err != nil {
		return vars, err
	}

	for _, match := range matches {
		var_name := regex.FindString(match)
		vars = append(vars, var_name)
	}

	return vars, nil
}
