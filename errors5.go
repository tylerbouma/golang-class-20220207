package main

import (
	"errors"
	"fmt"
)

func obiwan(name string) (string, error) {
	if name == "" {
		return "", errors.New("no name provided")
	}
	return (name + ", now that is a name I haven't heard in a long time"), nil
}

func main() {
	var name string

	fmt.Print("What was the name? (Just press enter)")
	fmt.Scanf("%s", &name)

	_, err := obiwan(name)

	if err != nil {
		fmt.Println("Could not run the obiwan function:", err)
		return
	}
}
