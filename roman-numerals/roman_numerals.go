// Package romannumerals contains a single function for working with Roman Numerals.
package romannumerals

import (
    "errors"
	"strconv"
)

type romanPair struct {
    one string
    five string
}

var romanNumerals = [4]romanPair{
    romanPair{"I", "V"},
    romanPair{"X", "L"},
    romanPair{"C", "D"},
    romanPair{"M", "_V"},
}

// ToRomanNumeral converts an integer value between 1 and 3000 into a string of roman numerals.
// This was meant to be so much simpler than it is.
func ToRomanNumeral(number int) (string, error) {
	
	// Roman's didn't count above 3000 and has no concept of 0.
    if number <= 0 || number > 3000   {
        return "", errors.New("input value was less than 1 or greater than 3000")
    }

	rn := ""
	
	// Convert the number into a string so we can access individual digits.
	numString := strconv.Itoa(number)
	
	// Loop through each individual digit.
	for index, char := range numString {

		// Whacky was to turn it back into an integer because it comes out as a Rune.
		digit, _ := strconv.Atoi(string(char))
		
		// Invert the index as we're working from biggest to smallest.
		i := len(numString) - 1 - index

		// Logic flow based on value of indiviual digits.
		if digit > 0 && digit < 4 {
            for j := 0; j < digit; j++ {
                rn += romanNumerals[i].one
			}
        } else if digit == 4 {
            rn += romanNumerals[i].one + romanNumerals[i].five
        } else if digit > 4 && digit < 9 {
            rn += romanNumerals[i].five
            for j := 0; j < (digit - 5); j++ {
                rn += romanNumerals[i].one
            }
        } else if digit == 9 {
            rn += romanNumerals[i].one + romanNumerals[(i + 1)].one
        } else {
            // 0
        }
	}

    return rn, nil
}