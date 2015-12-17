// Create a custom error type called appError that contains three fields, err error,
// message string and code int. Implement the error interface providing your own message
// using these three fields. Implement a second method named Temporary that returns false
// when the value of the code field is 9. Write a function called checkFlag that accepts
// a bool value. If the value is false, return a pointer of your custom error type
// initialized as you like. If the value is true, return a default error. Write a main
// function to call the checkFlag function and check the error.
package main

// Add imports.
import "errors"
import "fmt"

// Declare a struct type named appError with three fields, err of type error,
// message of type string and code of type int.
type appError struct {
	err     error
	message string
	code    int
}

// Declare a method for the appError struct type that implements the
// error interface.
func (e appError) Error() string {
	return fmt.Sprintf("App Error: %s, Message: %s, Code: %d", e.err, e.message, e.code)
}

// Declare a method for the appError struct type named Temporary that returns
// true when the value of the code field is 9.
func (e appError) Temporary() bool {
	return e.code != 9
}

// Declare the temporary interface type with a method named Temporary that
// takes no parameters and returns a bool.
type temporary interface {
	Temporary() bool
}

// Declare a function named checkFlag that accepts a boolean value and
// returns an error interface value.
func checkFlag(b bool) error {
	// If the parameter is false return an appError.
	errorMessage := fmt.Sprint("Flag is ", b)
	if b == false {
		return &appError{
			errors.New(errorMessage),
			"app error message string",
			8,
		}
	}
	// Return a default error.
	return errors.New(errorMessage)
}

// main is the entry point for the application.
func main() {
	// Call the checkFlag function to simulate an error of the
	// concrete type.
	err := checkFlag(true)

	// Check the concrete type and handle appropriately.
	switch e := err.(type) {
	//switch e := err {
	// e := err
	// switch e {
	//switch err {
	// Apply the case for the existence of the Temporary behavior.
	// Log the error and write a second message only if the
	// error is not temporary.
	case *appError:
		if e.Temporary() {
			fmt.Printf("Temp APP ERROR. Error value: %+v\n", e)
			return
		}
		fmt.Printf("APP ERROR. Error value: %+v\n", e)
		return
		// case *errorString:
		// 	if !e.Temporary() {
		// 		println("Temp ERROR STRING")
		// 	}
	// Apply the default case and just log the error.
	default:
		fmt.Printf("Logging DEFAULT error: %+v, %+v\n", e, err)
	}
}
