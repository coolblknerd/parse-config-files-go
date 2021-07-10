package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type yamlConfig struct {
	Enabled bool   `yaml:"enabled"`
	Path    string `yaml:"path"`
}

func main() {
	conf := yamlConfig{}
	file, _ := ioutil.ReadFile("config.yaml")

	err := yaml.Unmarshal(file, &conf)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println(conf.Enabled)
}
