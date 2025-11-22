package isogram

import "unicode"

// IsIsogram examines a word and assert if it is an isogram
func IsIsogram(word string) bool {

	characters := make(map[rune]int)

	for _, character := range word {
		if unicode.IsLetter(character) {
			if characters[unicode.ToUpper(character)] > 0 {
				return false
			} else {
				characters[unicode.ToUpper(character)]++
			}
		}
	}

	return true
}
