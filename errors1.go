package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("we don't have the power")
	fmt.Println("scooty says:", err)
}
