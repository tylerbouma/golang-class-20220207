package main

import "fmt"

type astro struct {
	name     string
	age      int
	mission  string
	isneeded bool
}

func (a astro) honorific() string {
	return "space hero: " + a.name
}

func (a astro) customhonorific(honor string) string {
	return honor + a.name
}

func (a *astro) suitup() {
	a.isneeded = true
}
func main() {
	ast1 := astro{"megan mcarthur", 53, "ISS", false}
	fmt.Println(ast1)

	fmt.Println(ast1.honorific())

	fmt.Println(ast1.customhonorific("Awesome Space Hero "))

	ast1p := &ast1
	fmt.Println(ast1p.customhonorific("awesome space hero "))

	fmt.Println("False to True")
	fmt.Println(ast1)
	ast1.suitup()
	fmt.Println(ast1)
}
