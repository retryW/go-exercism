// Package leap contains various functions for working with leap years.
package leap

// IsLeapYear returns a boolean value based on whether the provided year is a leap year.
func IsLeapYear(year int) bool {
	// Leap year is evenly divisible by 4, but not 100, except when also by 400.
	return year % 4 == 0 && !(year % 100 == 0 && year % 400 != 0)
}
