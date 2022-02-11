package main

import (
	"errors"
	"fmt"
)

func endit() error {
	return errors.New("we don't have the power")
}

func main() {
	err := endit()

	if err != nil {
		fmt.Println("we detected an error:", err)
		return
	}

	fmt.Println("release the hounds")
}
