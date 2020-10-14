// Package space contains helper functions relating to space and our solar system.
package space

type Planet string

// For converting seconds to an Earth calendar year.
var secToYear = 31557600.0

// Age returns the age of someone based on the calendar year equivalent for the input planet.
func Age(seconds float64, planet Planet) float64 {
	var orbitPeriod float64
	switch planet {
		case "Mercury":
			orbitPeriod = 0.2408467
		case "Venus":
			orbitPeriod = 0.61519726
		case "Earth":
			orbitPeriod = 1.0
		case "Mars":
			orbitPeriod = 1.8808158
		case "Jupiter":
			orbitPeriod = 11.862615	
		case "Saturn":
			orbitPeriod = 29.447498
		case "Uranus":
			orbitPeriod = 84.016846
		case "Neptune":
			orbitPeriod = 164.79132
		default:
			// Invalid input
			return 0
	}
	// Convert seconds to Earth years, and divide by the orbit period.
	return (seconds / secToYear) / orbitPeriod
}
