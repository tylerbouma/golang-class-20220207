package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("Enter text: ")
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("an error has occured while reading input. try again", err)
		return
	}

	input = strings.TrimSuffix(input, "\n")
	fmt.Println(input)
}
