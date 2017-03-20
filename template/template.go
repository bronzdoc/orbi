package template

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type template struct {
	Context  string
	Skeleton []map[string]string
}

func New(filepath string) *template {
	t := template{}
	data, err := ioutil.ReadFile(filepath)

	if err != nil {
		log.Fatal(err)
		return &template{}
	}

	if err = yaml.Unmarshal(data, &t); err != nil {
		log.Fatal(err)
		return &template{}
	}

	return &t
}
