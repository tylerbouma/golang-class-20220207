package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	drive := 0
	fmt.Println("get a mulligan on any drive under 60 yds")

	for drive <= 60 {
		rand.Seed(time.Now().UnixNano())
		drive = rand.Intn(251)
		fmt.Println("SWING!")
	}
	fmt.Println("your longest drive was", drive, "yards")
}
