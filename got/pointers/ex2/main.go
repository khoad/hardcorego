// Declare a struct type and create a value of this type. Declare a function
// that can change the value of some field in this struct type. Display the
// value before and after the call to your function.
package main

// Add imports.

// Declare a type named user.
type user struct {
	name string
	age  int
}

// Create a function that changes the value of one of the user fields.
func changeAge(u *user, age int) {
	// Use the pointer to change the value that the
	// pointer points to.
	u.age = age
}

// main is the entry point for the application.
func main() {
	// Create a variable of type user and initialize each field.
	khoa := user{
		name: "khoa",
		age:  25,
	}

	// Display the value of the variable.
	println("Before:", khoa.age)

	// Share the variable with the function you declared below.
	changeAge(&khoa, 10)

	// Display the value of the variable.
	println("After:", khoa.age)
}
