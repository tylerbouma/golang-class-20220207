package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	os.Setenv("ALTA", "3")
	fmt.Println("ALTA", os.Getenv("ALTA"))

	fmt.Println("RESEARCH:", os.Getenv("RESEARCH"))

	fmt.Println()

	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0])
	}
}
