// Package scrabble contains a single function to calculate the scrabble score of any given word.
package scrabble

import "unicode"

var scores = map[rune]int {
	'A': 1,
	'E': 1,
	'I': 1,
	'O': 1,
	'U': 1,
	'L': 1,
	'N': 1,
	'R': 1,
	'S': 1,
	'T': 1,
	'D': 2,
	'G': 2,
	'B': 3,
	'C': 3,
	'M': 3,
	'P': 3,
	'F': 4,
	'H': 4,
	'V': 4,
	'W': 4,
	'Y': 4,
	'K': 5,
	'J': 8,
	'X': 8,
	'Q': 10,
	'Z': 10,
}

// Score returns the score any word is worth in the game Scrabble.
func Score(word string) int {

	wordScore := 0

	// Loop over each character in the word.
	for _, char := range word {
		// Make it uppercase so I don't need double declaration in the map.
		char = unicode.ToUpper(char)
		// If the character exists as a key in the map, add it's value to the score.
		if value, keyExists := scores[char]; keyExists {
			wordScore += value;
		}
	}

	return wordScore
}
