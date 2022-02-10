package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var tries int

	fmt.Println("game: guess a number")
	fmt.Print("how many tries would you like to have?: ")
	fmt.Scan(&tries)
	fmt.Printf("you have %v tries\n", tries)

	source := rand.NewSource(time.Now().UnixNano())

	randomizer := rand.New(source)
	var secretNumber int
	if tries < 5 {
		secretNumber = randomizer.Intn(10)
		fmt.Println("guess a number between 0 and 10")
	} else {
		secretNumber = randomizer.Intn(100)
		fmt.Println("guess a number between 0 and 100")
	}

	var guess int

	for try := 1; try <= tries; try++ {
		fmt.Println("Round:", try)
		fmt.Println("Please enter your number")
		fmt.Scan(&guess)

		if guess < secretNumber {
			fmt.Println("sorry, your guess was too low")
		} else if guess > secretNumber {
			fmt.Println("sorry, your guess was too high")
		} else {
			fmt.Println("you win!")
			break
		}

		if try == tries {
			fmt.Println("game over!")
			fmt.Println("the correct number was:", secretNumber)
			break
		}
	}

	fmt.Println("thanks for playing")
}
