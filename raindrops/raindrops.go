// Package raindrops contains a single function useful for raindrop conversions.
package raindrops

import "fmt"

// Convert returns the sounds a given raindrop will make given a factor of 3, 5, or 7.
func Convert(rain int) string {
	
	raindrop := ""
	
	if rain % 3 == 0 {
		raindrop += "Pling"
	}
	if rain % 5 == 0 {
		raindrop += "Plang"
	}
	if rain % 7 == 0 {
		raindrop += "Plong"
	}

	// If none of the above were factors, write input value as output string.
	if raindrop == "" {
		return fmt.Sprint(rain)
	} else {
		return raindrop
	}
}