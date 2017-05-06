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

func (t *Template) Vars() map[string]string {
	return t.vars
}

func (t *Template) validateVars() error {
	varNames, err := findVars(string(t.content))
	if err != nil {
		return fmt.Errorf("findVars: %s", err)
	}

	var missingVars []string
	for _, name := range varNames {
		key := name
		if _, ok := t.vars[key]; !ok {
			missingVars = append(missingVars, key)
		}
	}

	if len(missingVars) > 0 {
		var errMsg bytes.Buffer
		for _, missingVar := range missingVars {
			errMsg.WriteString("  {{" + "." + missingVar + "}}\n")
		}

		return fmt.Errorf(
			"%s: missing vars:\n%s",
			t.name,
			errMsg.String(),
		)
	}
	return nil
}

func findVars(templateContent string) ([]string, error) {
	var vars []string

	regex, err := regexp.Compile(`{[^}]*}`)
	if err != nil {
		return vars, fmt.Errorf("regexp.Compile %s", err)
	}

	matches := regex.FindAllString(templateContent, -1)

	regex, err = regexp.Compile(`[a-zA-Z0-9_-]+`)
	if err != nil {
		return vars, fmt.Errorf("regexp.Compile %s", err)
	}

	for _, match := range matches {
		varName := regex.FindString(match)
		vars = append(vars, varName)
	}

	return vars, nil
}
