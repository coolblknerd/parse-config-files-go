package main

import (
	"fmt"

	"gopkg.in/gcfg.v1"
)

type iniConfig struct {
	Section struct {
		Enabled bool
		Path    string
	}
}

func main() {
	conf := iniConfig{}
	err := gcfg.ReadFileInto(&conf, "config.ini")
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println("Path: ", conf.Section.Path)
}
