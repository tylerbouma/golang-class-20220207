package main

import (
	"os"
	"text/template"
)

type Todo struct {
	Name        string
	Description string
}

func main() {
	td := Todo{"Test templates", "Let's test a template to see the magic."}

	t, err := template.New("todos").Parse("you have a task named \"{{.Name}}\" with description: \"{{.Description}}\"")

	if err != nil {
		panic(err)
	}

	err = t.Execute(os.Stdout, td)
	if err != nil {
		panic(err)
	}

	tbNew := Todo{"make breakfast", "grind and brew coffee"}

	err = t.Execute(os.Stdout, tbNew)

	if err != nil {
		panic(err)
	}
}
