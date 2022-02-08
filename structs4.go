package main

import "fmt"

type Astro struct {
	name        string
	age         int
	lastmission string
	isneeded    bool
}

type nasaMission struct {
	people  []Astro
	number  int
	message string
}

func main() {

	astro1 := Astro{"bob", 30, "mars", false}
	astro2 := Astro{"alice", 55, "pluto", true}
	astro3 := Astro{"rodrick", 22, "earth", true}

	var sl1 = []Astro{astro1, astro2, astro3}

	fmt.Println(astro1)
	fmt.Println(astro2)
	fmt.Println(astro3)

	fmt.Println(sl1)

	s := nasaMission{sl1, 3, "success"}
	fmt.Println(s)
	fmt.Printf("%+v", s)

}
