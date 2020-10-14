// Package hamming contains functions for working with DNA data.
package hamming

import "errors"

// Distance returns the Hamming Distance between two DNA strands.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		// Hamming distanced is only defined for sequences of equal length.
		return -1, errors.New("DNA length mismatch.")
	}

	ham := 0

	// Incrememt Hamming distance for each mismatch.
	for i:= 0; i < len(a); i++ {
		if a[i] != b[i] {
			ham++
		}
	}

	return ham, nil
}
