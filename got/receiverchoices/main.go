// Declare a struct type named Point with two fields, X and Y of type int.
// Implement a factory function for this type and a method that accepts
// this type and calculates the distance between the two points. What is
// the nature of this type?
package main

import "math"

// Add imports.
import "fmt"

// Point ...
// Declare struct type named Point that represents a point in space.
type Point struct {
	X int
	Y int
}

// New ...
// Declare a function named New that returns a Point based on X and Y
// positions on a graph.
func New(x, y int) Point {
	return Point{x, y}
}

// Distance ...
// Declare a method named Distance that finds the length of the hypotenuse
// between two points. Pass one point in and return the answer.
// Formula is the square root of (x2 - x1)^2 + (y2 - y1)^2
// Use the math.Pow and math.Sqrt functions.
func (a Point) Distance(b Point) float64 {
	return math.Sqrt(math.Pow(float64(a.X-b.X), 2) + math.Pow(float64(a.Y-b.Y), 2))
}

// main is the entry point for the application.
func main() {
	// Declare the first point.
	a := New(46, 32)

	// Declare the second point.
	b := Point{108, 54}

	// Calculate the distance and display the result.
	dist := a.Distance(b)
	fmt.Println("Distance", dist)
}
