package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("game: guess a number between 0 and 10")
	fmt.Println("you have three tries")

	source := rand.NewSource(time.Now().UnixNano())

	randomizer := rand.New(source)
	secretNumber := randomizer.Intn(10)

	var guess int

	for try := 1; try <= 3; try++ {
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

		if try == 3 {
			fmt.Println("game over!")
			fmt.Println("the correct number was:", secretNumber)
			break
		}
	}

	fmt.Println("thanks for playing")
}
