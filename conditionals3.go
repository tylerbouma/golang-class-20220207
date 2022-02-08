package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	if growth := rand.Intn(21 - -100) + -100; growth < 0 {
		fmt.Println("Looks like growth was neg:", growth)
	} else if growth < 10 {
		fmt.Println("Growth was a bit flat", growth)
	} else {
		fmt.Println("double digit growth:", growth)
	}
}
