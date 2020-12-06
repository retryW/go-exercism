// Package reverse contains a single function that returns a reversed string.
package reverse

// Reverse returns the input string but with all characters in the reverse order.
func Reverse(s string) string {
	// Convert string to a slice of runes.
	runeSlc := []rune(s)
	
	// Create empty slice for characters to be placed into in reverse.
	rev := []rune{}
	
	// Loop through the rune slice backwards and append each character to rev.
	for i := len(runeSlc) - 1; i >= 0; i-- {
		rev = append(rev, runeSlc[i])
	}
	return string(rev)
}