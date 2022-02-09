package main

import "fmt"

func main() {
	a := []string{"Alta3", "Research", "Loves", "Golang"}

	for i, s := range a {
		fmt.Println("position:", i, "value:", s)
	}

	for i := range a {
		fmt.Println("position:", i)
	}

	for pos := range a {
		fmt.Println("position:", pos)
	}
}
