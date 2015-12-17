// Declare a struct that represents a baseball player. Include name, atBats and hits.
// Declare a method that calculates a players batting average. The formula is Hits / AtBats.
// Declare a slice of this type and initialize the slice with several players. Iterate over
// the slice displaying the players name and batting average.
package main

import "fmt"

// Add imports.

// Declare a struct that represents a ball player.
type player struct {
	name   string
	atBats int
	hits   int
}

// Include field called name, atBats and hits.

// Declare a method that calculates the batting average for a player.
func (p *player) average() float32 {
	return float32(p.hits) / float32(p.atBats)
}

// main is the entry point for the application.
func main() {
	// Create a slice of players and populate each player
	// with field values.
	players := []player{
		{"bill", 4, 2},
		{"khoa", 10, 2},
	}

	// Display the batting average for each player in the slice.
	for _, player := range players {
		//println("Name:", player.name, "Average:", player.average())
		fmt.Printf("%s: AVG[%f]\n", player.name, player.average())
		fmt.Printf("%s: AVG[.%.f]\n", player.name, player.average()*1000)
	}
}
