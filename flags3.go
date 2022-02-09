package main

import (
	"flag"
	"fmt"
)

func main() {
	wordPtr := flag.String("word", "dont panic", "a string")
	numbPtr := flag.Int("numb", 42, "an int")
	boolPtr := flag.Bool("truthy", false, "a bool")

	var spiderman string
	flag.StringVar(&spiderman, "spiderman", "does whatever a spider can", "catch phrase")

	flag.Parse()

	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("truthy:", *boolPtr)
	fmt.Println("spiderman:", spiderman)
	fmt.Println("tail:", flag.Args())
}
