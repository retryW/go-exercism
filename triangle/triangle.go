// Package triangle provides funtions for working with triangle shapes.
package triangle

import "math"

type Kind int

const (
    // Create a sudo enum for the various types of triangles.
    NaT = 0  // not a triangle
    Equ = 1  // equilateral
    Iso = 2  // isosceles
    Sca = 3  // scalene
)

// KindFromSides evaluates the length of a triangles sides to determine what kind it is.
func KindFromSides(a, b, c float64) Kind {
	
	var k Kind

	switch {
		// If any of the attributes are not a number or is an , this is not a triangle, unless it's made of string...
		case math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c):
			k = NaT

		// If any of the attributes are infinity or negative infinity, this is not a triangle.
		case math.IsInf(a, 1) || math.IsInf(b, 1) || math.IsInf(c, 1) || math.IsInf(a, -1) || math.IsInf(b, -1) || math.IsInf(c, -1):
			k = NaT
		
		// If any sides are not greater than 0, this is not a triangle.
		case a <= 0 || b <= 0 || c <= 0:
			k = NaT

		// If the sum of the length of any two sides is not >= the third, this is not a triangle.
		case a + b < c || a + c < b || b + c < a:
			k = NaT

		// All sides are equal.
		case a == b && a == c:
			k = Equ

		// Only two sides are equal - Exercise says at least two, but I assume that's to simplify logic.
		case (a == b && a != c) || (a == c && a != b) || (b == c && a != b):
			k = Iso

		// No sides are equal.
		case a != b && a != c:
			k = Sca
	}

	return k
}
