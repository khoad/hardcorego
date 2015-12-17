// Write a problem that uses a buffered channel to maintain a buffer
// of four strings. In main, send the strings 'A', 'B', 'C' and 'D'
// into the channel. Then create 20 goroutines that receive a string
// from the channel, display the value and then send the string back
// into the channel. Once each goroutine is done performing that task,
// allow the goroutine to terminate.
package main

// Add Imports.
import "sync"

// Declare a constant and set the value for the number of goroutines.
const totalGr = 20

// Declare a constant and set the value for the number of buffers.
const totalBuffers = 4

// Declare a wait group.
var wg sync.WaitGroup

// Declare a buffered channel of type string and initialize it based on
// the constant you declared above.
var channel = make(chan string, totalBuffers)

// Declare a function for the goroutine that will process work
// from the buffered channel.
func worker(worker int) {
	// Receive a string from the channel.
	work := <-channel

	// Display the string.
	println("Worker", worker, "received", work)

	// Send the string back into the channel.
	channel <- work

	// Tell main this goroutine is done.
	wg.Done()
}

// main is the entry point for the application.
func main() {
	// Increment the wait group for the number of
	// goroutines based on the value of the constant.
	wg.Add(totalGr)

	// Create the number of goroutines based on the
	// value of the constant.
	for i := 1; i <= totalGr; i++ {
		go worker(i)
	}

	// Add strings in the buffered channel.
	for rune := 'A'; rune < 'A'+totalBuffers; rune++ {
		channel <- string(rune)
	}

	// Wait for all the work to get done.
	wg.Wait()
}
