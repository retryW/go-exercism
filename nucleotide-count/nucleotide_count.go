// Package dna contains a single function to assist in working with DNA.
package dna

import "errors"

// Histogram is a mapping from nucleotide to its count in given DNA.
type Histogram map[rune]int

// DNA is a list of nucleotides.
type DNA string

// Counts returns the number of each nucleotide found in the specified DNA input.
func (d DNA) Counts() (Histogram, error) {

	h := Histogram{'A':0, 'C':0, 'G':0, 'T':0}

	// Only process if the input actually contains a value.
	if len(d) > 0 {
		for _, char := range d {
			// Return an error if invalid input. 
			// I feel like this should be cleaner and simpler without requiring a new function or importing a stdlib package...
			if char != 'A' && char != 'C' && char != 'G' && char != 'T' {
				return nil, errors.New("input character is not a valid nucleotide")
			}

			// Update count of current nucleotide.
			h[char] += 1
		}
	}

	return h, nil
}
