// Package proverb provides a function for generating an ancient proverb.
package proverb

import "fmt"

// Proverb generates an ancient proverb based off the input parameters.
func Proverb(rhyme []string) []string {
	
	length := len(rhyme)

	// If we receive an empty slice, send one back.
	if length == 0 {
		return []string{}
	}

	// Create an empty slice with max length of the rhyme's length.
	prov := make([]string, length)

	// Loop over input values to generate the proverb.
	for i := 0; i < length - 1; i++ {
		prov[i] = fmt.Sprintf("For want of a %v the %v was lost.", rhyme[i], rhyme[i + 1])
	}

	prov[length - 1] = fmt.Sprintf("And all for the want of a %v.", rhyme[0])
	return prov
}
