package main

import (
	"fmt"
)

type Animal interface {
	Sound() string
	MakeSound()
}

type abstractAnimal struct{ Animal }

func (a abstractAnimal) MakeSound() {
	fmt.Printf("%v", a.Sound())
}

type Chicken struct{ abstractAnimal }

func (c Chicken) Sound() string {
	return "bwaak bwak bwak bwak\n"
}

func NewChicken() *Chicken {
	chicken := Chicken{abstractAnimal{}}
	chicken.abstractAnimal.Animal = chicken
	return &chicken
}

type Lion struct{ abstractAnimal }

func (d Lion) Sound() string {
	return "RAWWWR\n"
}

func NewLion() *Lion {
	lion := Lion{abstractAnimal{}}
	lion.abstractAnimal.Animal = lion
	return &lion
}

func main() {
	c := NewChicken()
	c.MakeSound()
	d := NewLion()
	d.MakeSound()
}
