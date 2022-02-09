package main

import (
	"html/template"
	"io"
	"log"
	"os"
)

func parse(path string) {
	t, err := template.ParseFiles(path)
	if err != nil {
		log.Print(err)
		return
	}

	f, err := os.Create(path)
	if err != nil {
		log.Println("create file: ", err)
		return
	}

	config := map[string]string{
		"textColor":      "#abcdef",
		"linkColorHover": "#ffaacc",
	}

	err = t.Execute(f, config)
	if err != nil {
		log.Print("execute: ", err)
		return
	}
	f.Close()
}

func main() {
	f, _ := os.Create("example.css")

	f.Write([]byte("the text color is {{.textColor}} and the link color is {{.linkColorHover}}"))

	f.Close()

	parse("example.css")

	f, _ = os.Open("example.css")
	io.Copy(os.Stdout, f)
	f.Close()
}
