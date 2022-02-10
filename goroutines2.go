package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func countDown(s int) {
	defer wg.Done()
	for i := s; i > 0; i-- {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

func main() {
	fmt.Println("NASA launch sequence starting")

	wg.Add(1)

	go countDown(10)

	fmt.Println("Launch sequence starting:")

	time.Sleep(time.Second)
	fmt.Println("Hey wait a second...")

	time.Sleep(time.Second)
	fmt.Println("I forgot my wallet!")

	// wait for the waitGroup queue to reach 0
	wg.Wait()

	fmt.Println("TAKE OFF!")
}
