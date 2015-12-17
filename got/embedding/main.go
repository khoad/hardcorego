// Copy the code from the template. Declare a new type called hockey
// which embeds the sports type. Implement the matcher interface for hockey.
// When implementing the match method for hockey, call into the match method
// for the embedded sport type to check the embedded fields first. Then create
// two hockey values inside the slice of matchers and perform the search.
package main

import (
	"fmt"
	"strings"
)

// matcher defines the behavior required for performing matching.
type matcher interface {
	match(searchTerm string) bool
}

// sport represents a sports team.
type sport struct {
	team string
	city string
}

// match checks the value for the specified term.
func (s sport) match(searchTerm string) bool {
	return strings.Contains(s.team, searchTerm) || strings.Contains(s.city, searchTerm)
}

// Declare a struct type named hockey that represents specific
// hockey information. Have it embed the sport type first.
type hockey struct {
	sport
	country string
}

// match checks the value for the specified term.
func (h hockey) match(searchTerm string) bool {
	// Make sure you call into match method for the embedded sport type.
	// Implement the search for the new fields.
	return h.sport.match(searchTerm) || strings.Contains(h.country, searchTerm)
}

// main is the entry point for the application.
func main() {
	// Define the term to match.
	term := "Miami"

	// Create a slice of matcher values to match.
	matchers := []matcher{
		hockey{sport{"Panthers", "Miami"}, "USA"},
		hockey{sport{"Canadiens", "Montreal"}, "Canada"},
	}

	// Display what we are trying to match.
	println("Matching for:", term)

	// Range of each matcher value and check the term.
	for _, m := range matchers {
		if m.match(term) {
			fmt.Println("FOUND:", m) // this is equivilent to "%v"?
			fmt.Printf("FOUND: %v\n", m)
			fmt.Printf("FOUND: %+v\n", m)
		}
	}
}
