package main

import (
	"fmt"
	"runtime"
	"strings"
)

func main() {
	gover := runtime.Version()

	switch {
	case strings.Contains(gover, "go1.17"):
		fmt.Println("This is the latest version", gover)
	case strings.Contains(gover, "go1.16"), strings.Contains(gover, "go1.15"):
		fmt.Println("older, but still good version: ", gover)
	default:
		fmt.Println("please upgrade")
	}
}
