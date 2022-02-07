package main

import "fmt"

var palm, python, tiger bool

var rock, scissors int = 42, 55

func main() {
	var i int

	var sun, moon = "eclipse", "waxing"

	fmt.Println(i, palm, python, tiger)

	fmt.Println("The value of the var rock is:", rock)
	fmt.Println("The value of the var scissors is:", scissors)

	fmt.Println("Look at the moon when it is", moon + ".")
	fmt.Println("A total", sun, "of the heart")
}
