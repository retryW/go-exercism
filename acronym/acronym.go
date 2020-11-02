// Package acronym contains a single function to help work with acronyms.
package acronym

import "strings"

var replacer = strings.NewReplacer("_", " ", "-", " ")

// Abbreviate returns the abbreviation of a common string.
func Abbreviate(s string) string {
	acr := ""

	// Replace any additional characters with whitespace.
	s = replacer.Replace(s)
	
	// Loop through each word in the string and append the first letter as a captial letter to our acronym.
	for _, word := range strings.Split(s, " ") {
		if len(word) > 0 {
			acr += strings.ToUpper(string(word[0]))
		}
	}

	return acr
}
