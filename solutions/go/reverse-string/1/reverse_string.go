package reverse

import "strings"

// Reverse reverses a string
func Reverse(word string) string {

	var length = len(word)
	var reversedWord = make([]string, length)

	for i, character := range word {
		position := -(i) + (length - 1)
		reversedWord[position] = string(character)
	}

	return strings.Join(reversedWord, "")
}
