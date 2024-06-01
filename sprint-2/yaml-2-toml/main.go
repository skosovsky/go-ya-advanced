package main

import (
	"log"

	"github.com/pelletier/go-toml/v2"
	"gopkg.in/yaml.v3"
)

type Data struct {
	ID     int    `yaml:"id"     toml:"id"`
	Name   string `yaml:"name"   toml:"name"`
	Values []byte `yaml:"values" toml:"values"`
}

const yamlData = `
id: 101
name: Gopher
values:
- 11
- 22
- 33
`

func main() {
	var data Data

	err := yaml.Unmarshal([]byte(yamlData), &data)
	if err != nil {
		log.Fatal(err)
	}

	tomlData, err := toml.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(tomlData))
}
