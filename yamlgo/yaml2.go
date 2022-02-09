package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type User struct {
	Name       string
	Occupation string
}

func main() {
	yfile, err := ioutil.ReadFile("villans.yaml")

	if err != nil {
		log.Fatal(err)
	}

	data := make(map[string]User)

	err2 := yaml.Unmarshal(yfile, &data)

	if err2 != nil {
		log.Fatal(err2)
	}

	for k, v := range data {
		fmt.Printf("%s: %s\n", k, v)
	}
}
