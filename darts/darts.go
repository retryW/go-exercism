// package darts contains a single function for scoring a given dart throw.
package darts

import "math"

// Score returns the points earned for a single throw of a dart at a dart board.
func Score(x, y float64) int {
	s := 0
	// Get the distance from the center (0,0) using a^2 + b^2 = c^2
	cs := math.Pow(x, 2) + math.Pow(y, 2)
	c := math.Sqrt(cs)

	if c > 10 {
		// Missed the target.
	} else if c > 5 {
		// Outer circle.
		s = 1
	} else if c > 1 {
		// Middle circle
		s = 5
	} else {
		// Inner circle
		s = 10
	}
	return s
}