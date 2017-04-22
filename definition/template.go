package definition

import (
	"bytes"
	"fmt"
	"io"
	"regexp"
	"text/template"
)

type Template struct {
	name    string
	content []byte
	vars    map[string]string
}

func NewTemplate(name string, content []byte, vars map[string]string) *Template {
	return &Template{
		name:    name,
		content: content,
		vars:    vars,
	}
}

func (t *Template) Execute(w io.Writer) (*Template, error) {
	if err := t.validateVars(); err != nil {
		return t, fmt.Errorf("validateVars: %s", err)
	}

	tmpl, err := template.New(t.name).Option("missingkey=error").Parse(
		string(t.content),
	)
	if err != nil {
		return t, fmt.Errorf("text/template: %s", err)
	}

	err = tmpl.Execute(w, t.vars)
	if err != nil {
		return t, fmt.Errorf("text/template: %s", err)
	}

	return t, nil
}

func (t *Template) Name() string {
	return t.name
}

func (t *Template) Content() []byte {
	return t.content
}

func (t *Template) Vars(content []byte) map[string]string {
	return t.vars
}

func (t *Template) validateVars() error {
	var_names, err := findVars(string(t.content))
	if err != nil {
		return fmt.Errorf("findVars: %s", err)
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
			err_msg.WriteString("  {{" + "." + missing_var + "}}\n")
		}

		return fmt.Errorf(
			"%s: missing vars:\n%s",
			t.name,
			err_msg.String(),
		)
	}
	return nil
}

func findVars(template_content string) ([]string, error) {
	var vars []string

	regex, err := regexp.Compile(`{[^}]*}`)
	if err != nil {
		return vars, fmt.Errorf("regexp.Compile %s", err)
	}

	matches := regex.FindAllString(template_content, -1)

	regex, err = regexp.Compile(`[a-zA-Z0-9_-]+`)
	if err != nil {
		return vars, fmt.Errorf("regexp.Compile %s", err)
	}

	for _, match := range matches {
		var_name := regex.FindString(match)
		vars = append(vars, var_name)
	}

	return vars, nil
}
