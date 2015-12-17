// Declare a nil slice of integers. Create a loop that appends 10 values to the
// slice. Iterate over the slice and display each value.
//
// Declare a slice of five strings and initialize the slice with string literal
// values. Display all the elements. Take a slice of index one and two
// and display the index position and value of each element in the new slice.
package main

import "fmt"

// Add imports.

// main is the entry point for the application.
func main() {
	// Declare a nil slice of integers.
	var nums []int

	// Appends numbers to the slice.
	for index := 0; index < 10; index++ {
		nums = append(nums, index)
	}

	// Display each value in the slice.
	for _, num := range nums {
		println(num)
	}

	// Declare a slice of strings and populate the slice with names.
	names := []string{"khoa", "bob", "bill"}

	// Display each index position and slice value.
	for i, name := range names {
		println("Index", i, "Value", name)
	}

	// Take a slice of index 1 and 2 of the slice of strings.
	slicenames := names[1:3]

	// Display each index position and slice values for the new slice.
	for i, n := range slicenames {
		//println("Index", i, "Value", n)
		fmt.Printf("Index: %d Value: %s\n", i, n)
	}
}
