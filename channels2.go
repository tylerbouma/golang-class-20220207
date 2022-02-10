package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"sync"
)

var accountBalance = 0
var mutex = &sync.Mutex{}

func increaseBalance(amt int) {
	mutex.Lock()
	accountBalance += amt
	mutex.Unlock()
}

func reportAndExit(msg string) {
	fmt.Println(msg)
	os.Exit(-1)
}

func main() {
	if len(os.Args) < 2 {
		reportAndExit("\nUsage: go run channels02.go <number of updates per thread>")
	}
	iterations, err := strconv.Atoi(os.Args[1])
	if err != nil {
		reportAndExit("Bad command-line argument: " + os.Args[1])
	}

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < iterations; i++ {
			increaseBalance(1)
			fmt.Println("Adding")
			runtime.Gosched()
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < iterations; i++ {
			increaseBalance(-1)
			fmt.Println("Removing")
			runtime.Gosched()
		}
	}()

	wg.Wait()
	fmt.Println("Final balance: ", accountBalance)
}
