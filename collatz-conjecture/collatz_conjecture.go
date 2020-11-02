// Package collatzconjecture contains a single function useful for counting steps taken to complete the Collatz Conjecture.
package collatzconjecture

import "errors"

// CollatzConjecture calculates the steps taken for an integer to reach 1 using the 3n+1 Collatz Conjecture.
func CollatzConjecture (n int) (int, error) {
	count := 0

	// Return an error if input is not a positive integer.
	if n < 1 {
		return count, errors.New("input parameter was not a positive integer")
	}

	// Loop until we reach n = 1
	for n != 1 {
		if n % 2 == 0 { // Even number
			n = n / 2
		} else { // Odd number
			n = (3 * n) + 1
		}
		count++
	}

	return count, nil
}