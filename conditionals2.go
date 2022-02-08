package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var age int = 42

	fmt.Println("given that age is", age)

	if age%2 == 0 {
		fmt.Println("age is even")
	} else {
		fmt.Println("age is odd")
	}
}
