// Package strain contains functions useful for working with strains.
package strain

type Ints []int
type Lists [][]int
type Strings []string

// TODO: Maybe find a way to template these functions as they're essentially the same?
// It sure would be nice to have generics right about now...

// Keep (Ints) returns a slice of integers that pass the input predictate. 
func (ints Ints) Keep(f func(int) bool) Ints {
	var outInts Ints
	for i := 0; i < len(ints); i++ {
		if f(ints[i]) {
			outInts = append(outInts, ints[i])
		}
	}
	return outInts
}

// Discard (Ints) returns a slice of integers that fail the input predictate.
func (ints Ints) Discard(f func(int) bool) Ints {
	var outInts Ints
	for i := 0; i < len(ints); i++ {
		if !f(ints[i]) {
			outInts = append(outInts, ints[i])
		}
	}
	return outInts
}

// Keep (Lists) returns a slice of integer arrays that pass the input predictate.
func (l Lists) Keep(f func([]int) bool) Lists {
	var outLists Lists
	for i := 0; i < len(l); i++ {
		if f(l[i]) {
			outLists = append(outLists, l[i])
		}
	}
	return outLists
}

// Keep (Strings) returns a slice of strings that pass the input predictate.
func (s Strings) Keep(f func(string) bool) Strings {
	var outStrings Strings
	for i := 0; i < len(s); i++ {
		if f(s[i]) {
			outStrings = append(outStrings, s[i])
		}
	}
	return outStrings
}
