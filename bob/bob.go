// Package bob contains a single function to aid in speaking to your imaginary friend Bob.
package bob

import "regexp"

// Hey returns a piece of dialog from your imaginary friend Bob.
func Hey(remark string) string {
	var answer string
	// Pretty dirty regex, maybe using forward lookups is cleaner? :\
	switch {
	// No input or all whitespace.
	case remark == "" || MatchStringAsBool(`^\s+$`, remark):
		answer = "Fine. Be that way!"
	// All caps question.
	case MatchStringAsBool(`^[^a-z]*[A-Z]+[^a-z]*\?{1}\s*$`, remark):
		answer = "Calm down, I know what I'm doing!"
	// All caps statement.
	case MatchStringAsBool(`^[^a-z]*[A-Z]+[^a-z]*!*\s*$`, remark):
		answer = "Whoa, chill out!"
	// Regular question.
	case MatchStringAsBool(`\?{1}\s*$`, remark):
		answer = "Sure."
	default:
		answer = "Whatever."
	}
	return answer
}

// MatchStringAsBool returns a boolean value of whether a match was found given some regex and a string.
func MatchStringAsBool(r string, s string) bool {
	match, _ := regexp.MatchString(r, s)
	return match
}
