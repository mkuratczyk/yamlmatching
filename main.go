package main

import (
	"fmt"

	yaml "gopkg.in/yaml.v2"
)

type item struct {
	Name string `yaml:"name,omitempty"`
}

type collection struct {
	Items []item `json:"item,omitempty"`
}

func generate() string {
	myItems := collection{
		Items: []item{
			item{"item1"},
		},
	}

	jobsYAML, err := yaml.Marshal(myItems)
	if err != nil {
		return ""
	}
	return string(jobsYAML)
}

func main() {
	fmt.Println(generate())
}
