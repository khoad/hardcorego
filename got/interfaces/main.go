// Declare an interface named speaker with a method named speak. Declare a struct
// named english that represents a person who speaks english and declare a struct named
// chinese for someone who speaks chinese. Implement the speaker interface for each
// struct using a value receiver and these literal strings "Hello World" and "你好世界".
// Declare a variable of type speaker and assign the address of a value of type english
// and call the method. Do it again for a value of type chinese.
//
// Add a new function named sayHello that accepts a value of type speaker.
// Implement that function to call the speak method on the interface value. Then create
// new values of each type and use the function.
package main

// Add imports.

// Declare the speaker interface with a single method called speak.
type speaker interface {
	speak()
}

// Declare an empty struct type named english.
type english struct{}

// Declare a method named speak for the english type
// using a value receiver. "Hello World"
func (e english) speak() {
	println("Hello World")
}

// Declare an empty struct type named chinese.
type chinese struct{}

// Declare a method named speak for the chinese type
// using a value receiver. "你好世界"
func (c chinese) speak() {
	println("你好世界")
}

// sayHello accepts values of the speaker type.
func sayHello(s speaker) {
	// Call the speak method from the speaker parameter.
	s.speak()
}

// main is the entry point for the application.
func main() {
	// Declare a variable of the speaker type set to its zero value.
	var s speaker

	// Declare a variable of type english and assign it to
	// the speaker variable.
	var e english // using var to declare a zero value variable
	s = e
	// Call the speak method from the speaker variable.
	s.speak()

	// Declare a variable of type chinese and assign it to
	// the speaker variable.
	c := chinese{} // using empty litteral
	s = c
	// Call the speak method from the speaker variable.
	s.speak()

	// Call the sayHello function passing each concrete type.
	// all of these ways work
	sayHello(e)
	sayHello(c)
	sayHello(&e)
	sayHello(new(chinese))
	sayHello(chinese{})
}
