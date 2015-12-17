// Write a program where two goroutines pass an integer back and forth
// ten times. Display when each goroutine receives the integer. Increment
// the integer with each pass. Once the integer equals ten, terminate
// the program cleanly.
package main

// Add imports.
import (
	"fmt"
	"sync"
)

// Declare a wait group variable.
var wg sync.WaitGroup

// Declare a function for the goroutine. Pass in a name for the
// routine and a channel of type int used to share the value back and forth.
func goroutine(name string, ch chan int) {
	for {
		// Receive on the channel, waiting for the integer.
		num, ok := <-ch

		// Check if the channel is closed.
		if !ok {
			// Call done on the waitgroup.
			wg.Done()
			// Display the goroutine is finished and return.
			fmt.Printf("%v is finished with %v\n", name, num)
			return
		}

		// Display the integer value received.
		fmt.Printf("%v received %v\n", name, num)

		// Check is the value from the channel is 10.
		if num == 10 {
			// Close the channel.
			close(ch)
			// Call done on the waitgroup.
			wg.Done()
			// Display the goroutine is finished and return.
			fmt.Printf("%v is finished with %v\n", name, num)
			return
		}

		// Increment the value by one and send in back through
		// the channel.
		num++
		ch <- num
	}
}

// main is the entry point for the application.
func main() {
	// Declare and initialize an unbuffered channel
	// of type int.
	ch := make(chan int)

	// Increment the wait group for each goroutine
	// to be created.
	wg.Add(2)

	// Create the two goroutines.
	go goroutine("Bob", ch)
	go goroutine("Bill", ch)

	// Send the initial integer value into the channel.
	ch <- 1

	// Wait for all the goroutines to finish.
	wg.Wait()
}
