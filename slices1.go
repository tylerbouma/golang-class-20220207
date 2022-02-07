package main

import "fmt"

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	slice := array[:]

	slice[1] = 7
	fmt.Println("modifying slices")
	fmt.Println("Array =", array)
	fmt.Println("slice = ", slice)

	array[1] = 2
	fmt.Println("modifying underlying array")
	fmt.Println("array =", array)
	fmt.Println("slice =", slice)
}
