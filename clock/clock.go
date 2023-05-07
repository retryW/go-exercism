// Package clock contains functions to interact with a rudimentary 24 hour clock
package clock

// Import fmt for string manipulation
import "fmt"

// Clock is a type that tracks time based on total minutes
type Clock struct {
	minutes int
}

const maxMinutes = 1440

// Returns a new 24 hour clock with the given time
func New(h, m int) Clock {

	// Create a new clock at 00:00
	c := Clock{minutes: 0}

	// Calculate the total minutes to adjust the clock by, converting hours to minutes
	mins := m + h*60

	// If the current time is in the negative, subtract instead of add
	if mins < 0 {
		// Return the clock with the subtracted minutes
		return c.Subtract(mins)
	}

	// Return the clock with the added minutes
	return c.Add(mins)
}

// Add m minutes to the clock
func (c Clock) Add(m int) Clock {

	// Add minutes to clock
	c.minutes += m

	// Ensure time loops back to the start of the 24 hour period
	// Use maxMinutes-1 to ensure midnight is the same as zero
	if c.minutes > maxMinutes-1 {
		c.minutes %= maxMinutes
	}

	// Return the clock
	return c
}

// Subtract m minutes from the clock
func (c Clock) Subtract(m int) Clock {

	// Make sure the value is a positive integer
	if m < 0 {
		m *= -1
	}

	// Subtract the minutes from the clock
	c.minutes -= m

	// Ensure time loops back to the end of the 24 hour period
	if c.minutes < 0 {
		// Convert the current minutes to a positive number, then get the modulo, then subtract from midnight for the final time
		c.minutes = (maxMinutes - (c.minutes * -1 % maxMinutes))
	}

	// Return the clock
	return c
}

// Return a string of the clock's current time in 24 hour format HH:mm
func (c Clock) String() string {
	// Pad out string with 0's to 2 digits
	return fmt.Sprintf("%02d:%02d", (c.minutes/60)%24, c.minutes%60)
}
