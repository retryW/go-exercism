/*
Package gigasecond provides helper functions for working with gigaseconds.
A gigasecond is 10^9 (1,000,000,000) seconds.
*/
package gigasecond

// import path for the time package from the standard library.
import "time"

// AddGigasecond adds a gigasecond to the time passed in as a parameter.
func AddGigasecond(t time.Time) time.Time {
	// Convert a gigasecond into nanoseconds and add to time
	var ns time.Duration = 1000000000 * 1e+9
	return t.Add(ns)
}
