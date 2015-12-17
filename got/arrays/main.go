// Declare an array of 5 strings with each element initialized to its zero value.
//
// Declare a second array of 5 strings and initialize this array with literal string
// values. Assign the second array to the first and display the results of the first array.
// Display the string value and address of each element.
package main

// Add imports.

// main is the entry point for the application.
func main() {
	// Declare an array of 5 strings set to its zero value.
	var zeroStrings [5]string
	// Declare an array of 5 strings and pre-populate it with names.
	names := [5]string{"a", "b", "c", "d", "e"}
	// Assign the populated array to the array of zero values.
	zeroStrings = names

	// Iterate over the first array declared.
	// Display the string value and address of each element.
	for _, name := range zeroStrings {
		println(name, &name)
	}
}
