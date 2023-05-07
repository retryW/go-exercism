// Package luhn contains function for working with the luhn algorithm
package luhn

import (
	"strings"
	"unicode"
)

func Valid(id string) bool {

	// Remove all whitespace
	id = strings.ReplaceAll(id, " ", "")

	// Cast the byte array to a rune slice for easy index access
	idRuneSlice := []rune(id)

	irsLen := len(idRuneSlice)

	// String of length 1 or less are invalid
	if irsLen <= 1 {
		return false
	}

	// To track the total of all processed digits
	total := 0

	// To track alternating digit
	toggle := false

	// Process the input digits
	for i := irsLen - 1; i >= 0; i-- {

		// Non-digits are invalid
		if !unicode.IsDigit(idRuneSlice[i]) {
			return false
		}

		// Convert rune to int, this feels weird
		val := int(idRuneSlice[i] - '0')

		// Only double every second digit
		// If doubled digit is greater than 9, subtract 9
		if toggle {
			val *= 2
			if val > 9 {
				val -= 9
			}
		}

		// Add the doubled amount to the total of all digits
		total += val
		toggle = !toggle
	}

	// If the sum is evenly divisible by 10, then the number is valid
	return total%10 == 0
}
