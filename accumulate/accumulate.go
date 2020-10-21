// Package accumulate contains a single function to help work with collections.
package accumulate

type ToLower func (string) string

// Accumulate loops through an array of strings and performs an operation that was passed in as a variable.
func Accumulate(strArray []string, op ToLower) []string {
	
	if len(strArray) < 1 {
		return []string{}
	}

	for i := 0; i < len(strArray); i++ {
		strArray[i] = op(strArray[i])
	}
	return strArray
}
