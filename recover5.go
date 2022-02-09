package main

import "fmt"

func waterLawn() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("error detected and recovered in the watering system")
		}
	}()

	fmt.Println("begin watering system.")
	water(5)

	fmt.Println("watering system finished")
}

func water(g int) {
	fmt.Println("starting", g, "watering systems")

	for i := 1; i < g+1; i++ {
		fmt.Println("turning on sprinkler at station:", i)
		defer fmt.Println("turning off sprinkler at station:", i)

		if i == 3 {
			fmt.Println("panicking")
			panic(fmt.Sprintf("error occured on iteration %v with function input %v", i, g))
		}
	}
}

func main() {
	waterLawn()

	fmt.Println("begin measuring soil pH.")
	fmt.Println("begin measuring soil temp.")
	fmt.Println("completed gardening routine")
}
