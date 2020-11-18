// Package isogram contains a single function to aid in determining if a phrase is an isogram.
package isogram

import (
	"regexp"
	"unicode"
)

// IsIsogram determines if a word or phrase is an isogram (a "nonpattern word").
func IsIsogram(str string) bool {
	var used []string
	// Loop through each character in the string.
	for _, v := range str {
		// Make uppercase to half required checks, and convert to string for ease of use.
		s := string(unicode.ToUpper(v))
		// Ensure the character isn't a member of the used character list.
		for _, c := range used {
			if match, _ := regexp.MatchString(c, s); match {
				// Not an isogram.
				return false
			}
		}
		// If the character is a letter, add it to the used character list.
		if match, _ := regexp.MatchString(`[a-zA-Z]`, s); match {
			used = append(used, s)
		}
	}
	// Is an isogram.
	return true
}
