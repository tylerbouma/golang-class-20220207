package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type Record struct {
	Item string `yaml:"item"`
	Col  string `yaml:"color"`
	Size string `yaml:"postage"`
}

type Config struct {
	Record Record `yaml:"Settings"`
}

func main() {
	config := Config{Record: Record{Item: "window", Col: "blue", Size: "small flat"}}

	data, err := yaml.Marshal(&config)

	if err != nil {
		log.Fatal(err)
	}

	err2 := ioutil.WriteFile("config.yaml", data, 0644)

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println("data written")
}
