package main

import "fmt"

type Shark struct {
	name string
}

func (a Shark) Swim() {
	fmt.Println(a.name, "shark is swimming...")
}

func (a *Shark) SwimFaster() {
	fmt.Println(a.name, "shark is swimming...")
}

func main() {
	fish := Shark{"Tiger"}
	fish.Swim()
	fish.SwimFaster()
}
