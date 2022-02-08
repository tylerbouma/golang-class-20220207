package main

import (
	"errors"
	"fmt"
)

func rollchar(firstName string, lastName string) (string, error) {
	if lastName == "Turnip" || lastName == "Radish" || lastName == "Potato" {
		return firstName + " the " + lastName, errors.New("veggies are not a suitable name")
	}
	return firstName + " the " + lastName, nil
}

func main() {
	fmt.Println("Welcome to the Character Generator")

	playerChar, err := rollchar("Gandalf", "Turnip")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(playerChar, "generated")
	}
}
