package main

import (
	"errors"
	"fmt"
)

func obiwan(name string) (string, error) {
	if name == "" {
		return "", errors.New("no name provided")
	}

	return (name + ", now that is a name I haven't heard in a long time."), nil
}

func main() {
	var name string

	fmt.Print("what was the name?: a")
	fmt.Scanf("%s", &name)

	name, err := obiwan(name)

	if err != nil {
		fmt.Println("Could not run the obiwan function:", err)
		return
	}

	fmt.Println("Obi-Wan Says:", name)
}
