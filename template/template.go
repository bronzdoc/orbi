package template

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Template struct {
	Context  string
	Skeleton []map[string]interface{}
}

func New(filepath string) *Template {
	t := Template{}
	data, err := ioutil.ReadFile(filepath)

	if err != nil {
		log.Fatal(err)
	}

	if err = yaml.Unmarshal(data, &t); err != nil {
		log.Fatal(err)
		return &Template{}
	}

	return &t
}
