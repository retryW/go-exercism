// Package clock contains functions to interact with a rudimentary 24 hour clock
package clock

// Import fmt for string manipulation
import "fmt"

// A simple Clock type that only tracks hours and minutes
type Clock struct {
	hours   int
	minutes int
}

// Returns a new 24 hour clock with the given hour and minutes value
func New(h, m int) Clock {
	// Create a blank clock set to 00:00
	c := Clock{
		hours:   0,
		minutes: 0,
	}
	// Calculate the total number of minutes given via input
	min := (h * 60) + m
	if min >= 0 {
		// Total minutes was positive, so add to base clock
		return c.Add(min)
	} else {
		// Total minute was negative, so substract from base clock after converting to a positive integer.
		return c.Subtract(min * -1)
	}
}

// Add m minutes to the clock
func (c Clock) Add(m int) Clock {

	// Add total minutes to clock
	c.minutes += m

	// While minutes are greater than an hour, process the hour hand
	for c.minutes > 59 {

		// If the next hour is greater than 23 roll it over to 00, else add the hour as normal
		if c.hours+1 > 23 {
			c.hours = 0
		} else {
			c.hours++
		}

		// Remove the added hour from the minutes
		c.minutes -= 60
	}

	// Return the clock
	return c
}

// Subtract m minutes from the clock
func (c Clock) Subtract(m int) Clock {

	// Subtract the total minutes from the clock
	c.minutes -= m

	// While the minutes are negative, process hours
	for c.minutes < 0 {

		// If the previous is less than 0, roll it back to 23, else subtract an hour as normal
		if c.hours-1 < 0 {
			c.hours = 23
		} else {
			c.hours--
		}

		// Add the subtracted hour to the minutes
		c.minutes += 60
	}

	// Return the clock
	return c
}

// Return a string of the clock's current time in 24 hour format HH:mm
func (c Clock) String() string {
	// Pad out string with 0's to 2 digits
	return fmt.Sprintf("%02d:%02d", c.hours, c.minutes)
}
