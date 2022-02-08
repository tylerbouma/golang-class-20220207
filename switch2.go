package main

import (
	"fmt"
	"time"
)

func main() {
	watch := time.Now()

	switch {
	case watch.Hour() < 6:
		fmt.Println("Go back to sleep.")
	case watch.Hour() < 12:
		fmt.Println("Good morning")
	case watch.Hour() < 18:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("counting sheep... zzz")
	}
}
