package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	if 10%5 == 0 {
		fmt.Println("10 is divisible by 5")
	}

	if v := rand.Intn(101); v > 60 {
		fmt.Println("You tested in the top 60%. Your score was: ", v)
	}
	fmt.Println("We want to thank everyone for taking the test")

	fmt.Println("Enter your desired name: ")
	var first string
	fmt.Scanln(&first)

	if length := len(first); length > 11 {
		fmt.Println("Screen names have a limit of 18 characters \n only the first 11 will be displayed")
	}
}
