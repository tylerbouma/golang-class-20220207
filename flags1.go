package main

import (
	"flag"
	"fmt"
)

func main() {
	num := flag.Int("n", 5, "How angry is the hulk?")
	flag.Parse()

	n := *num

	for i := 0; i < n; i++ {
		fmt.Println("HULK SMASH!")
	}
}
