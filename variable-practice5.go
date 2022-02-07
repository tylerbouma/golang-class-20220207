package main

import (
	"fmt"
)

func main() {
	shadow := "world"
	fmt.Println(shadow)

	func() {
		shadow := "hello"
		fmt.Println(shadow)
	}()

	fmt.Println(shadow)
}
