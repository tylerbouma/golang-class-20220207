package main

import (
	"fmt"
	"os"
)

func main() {
	argsWithProg := os.Args

	argsWithoutProg := os.Args[1:]

	arg := os.Args[3]

	last_arg := os.Args[len(os.Args)-1]

	fmt.Println("all arguments:", argsWithProg)
	fmt.Println("all arguments, without program name:", argsWithoutProg)

	fmt.Println("3rd arg:", arg)

	fmt.Println("last arg:", last_arg)

	fmt.Println("number of args:", len(argsWithoutProg))

	for i, a := range argsWithoutProg {
		fmt.Println("arg ", a, "is at position ", i)
	}
}
