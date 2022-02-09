package main

import "fmt"

func main() {
	total := 0

	for alta := 0; alta < 4; alta++ {
		fmt.Println("The value of alta:", alta)

		total += alta
	}

	fmt.Println("\nthe value of total is:", total)
}
