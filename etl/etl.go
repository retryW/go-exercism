// Package etl contains a single function to aid with Extract-Transform-Load (ETL).
package etl

import "strings"

// Transform takes a scrabble score in legacy format and returns the same score in the new format.
func Transform(m map[int][]string) map[string]int {
	var out = map[string]int{}
	for index, value := range m {
		for i := 0; i < len(value); i++ {
			out[strings.ToLower(value[i])] = index
		}
	}
	
	return out
}