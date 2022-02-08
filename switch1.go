package main

import (
	"fmt"
	"runtime"
)

func main() {
	switch gover := runtime.Version(); gover {
	case "go1.17":
		fmt.Println("You are using the latest version of GoLang")
	case "go1.16", "go1.16.5", "go1.15":
		fmt.Println("You are using an older but acceptable version")
	default:
		fmt.Println("I cannot make a recommendation")
	}
}
