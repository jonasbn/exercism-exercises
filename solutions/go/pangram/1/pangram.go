package pangram

import (
	"unicode"
)

// IsPangram evaluated a string input and returns indication of whether a string is a pangram
func IsPangram(word string) bool {
	characters := make(map[rune]int)

	if word == "" {
		return false
	}

	for _, character := range word {
		if unicode.IsLetter(character) {
			characters[unicode.ToUpper(character)]++
		}
	}

	keys := []rune{}
	for key := range characters {
		keys = append(keys, key)
	}

	if len(keys) == 26 {
		return true
	}

	return false
}
