// Declare a struct type to maintain information about a user (name, email and age).
// Create a value of this type, initalize with values and display each field.
//
// Declare and initialize an anonymous struct type with the same three fields. Display the value.
package main

// Add imports.
import "fmt"

// Add user type and provide comment.
type user struct {
	name  string
	email string
	age   int
}

// main is the entry point for the application.
func main() {
	// Declare variable of type user and init using a struct literal.
	khoa := user{
		name:  "Khoa",
		email: "khoa.nguyen@nordstrom.com",
		age:   25,
	}
	// Display the field values.
	fmt.Println("Name", khoa.name)
	fmt.Println("Email", khoa.email)
	fmt.Println("Age", khoa.age)

	// Declare a variable using an anonymous struct.
	ed := struct {
		name  string
		email string
		age   int
	}{
		name:  "Ed",
		email: "ed@ardanstudio.com",
		age:   46,
	}
	// Display the field values.
	fmt.Printf("%+v\n", ed)

	fmt.Println("Name", ed.name)
	fmt.Println("Email", ed.email)
	fmt.Println("Age", ed.age)

	// e := struct {
	// 	name  string
	// 	email string
	// 	age   int
	// }{
	//
	// }
	khoa = ed

	fmt.Println("Name", khoa.name)
	fmt.Println("Email", khoa.email)
	fmt.Println("Age", khoa.age)
}
