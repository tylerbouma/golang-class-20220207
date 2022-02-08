package main

import "fmt"

type Coord struct {
	X int
	Y int
}

func main() {
	fmt.Println(Coord{1, 2})

	z := Coord{42, 100}
	z.Y = 180
	fmt.Println(z)
	fmt.Println("The X coordianate is:", z.X)
	fmt.Println("The Y coordianate is:", z.Y)
}
