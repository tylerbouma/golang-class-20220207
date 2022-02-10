/* Alta3 Reseach | RZFeeser
   Goroutine - With Mutex     */

package main

import (
	"fmt"
	"sync"
)

var balance int

func init() {
	balance = 1000
}

func deposit(value int, wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	fmt.Printf("Depositing %d to account with balance: %d\n", value, balance)
	balance += value
	m.Unlock()
	wg.Done()
}

func withdraw(value int, wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	fmt.Printf("Withdrawing %d from account with balance: %d\n", value, balance)
	balance -= value
	m.Unlock()
	wg.Done()
}

func main() {
	fmt.Println("Go 'Using the Mutex' Example")

	// had to pass the mutex into the functions to avoid it being ignored
	var m sync.Mutex
	var wg sync.WaitGroup

	wg.Add(3)
	go withdraw(750, &wg, &m)
	go deposit(1000, &wg, &m)
	go withdraw(500, &wg, &m)

	wg.Wait() // wait until we finish

	fmt.Printf("New Balance %d\n", balance)
}
