package main

import (
	"html/template"
	"os"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func main() {
	tmpl := template.Must(template.ParseFiles("index.html"))

	data := TodoPageData{
		PageTitle: "My TODO list",
		Todos: []Todo{
			{Title: "laundry", Done: false},
			{Title: "pull weeds in the garden", Done: true},
			{Title: "sweep the stairs", Done: true},
		},
	}
	tmpl.Execute(os.Stdout, data)
}
