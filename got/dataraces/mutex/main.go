// https://play.golang.org/p/r-lOOe5PbI

// Program for an exercise to fix a race condition.
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// mutex is used to define a critical section of code.
var mutex sync.Mutex

// numbers maintains a set of random numbers.
var numbers []int

// init is called prior to main.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// main is the entry point for all Go programs.
func main() {
	// wg is used to manage concurrency.
	var wg sync.WaitGroup
	wg.Add(3)

	// Create three goroutines to generate random numbers.
	for i := 0; i < 3; i++ {
		go func() {
			random(10)
			wg.Done()
		}()
	}

	// Wait for all the goroutines to finish.
	wg.Wait()

	// Display the set of random numbers.
	for i, number := range numbers {
		fmt.Println(i, number)
	}
}

// random generates random numbers and stores them into a slice.
func random(amount int) {
	// Generate as many random numbers as specified.
	for i := 0; i < amount; i++ {
		n := rand.Intn(100)
		mutex.Lock()
		{
			numbers = append(numbers, n)
		}
		mutex.Unlock()
	}
}
