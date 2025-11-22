package rotationalcipher

import (
	"unicode"
)

func RotationalCipher(plain string, shiftKey int) string {
	rotated_string := ""

	// Create a map of the alphabet
	alphabet := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	alphabet_length := len(alphabet)

	// Loop through the plain text, our original string
	for _, character := range plain {

		// Check if the value is a letter, non-letters will be ignored
		if unicode.IsLetter(character) {

			// Check if the letter is upper case, so we keep case
			isUpper := unicode.IsUpper(character)

			// Loop through the alphabet to find the position of the letter
			for index, letter := range alphabet {

				// Do our letter match a letter in the alphabet
				if unicode.ToLower(character) == letter {

					// Yes, we found a match, now we need to rotate the letter
					position := index + shiftKey

					var rotated_value rune

					// Do our position exceed the length of the alphabet
					if position >= alphabet_length {
						rotated_value = alphabet[(position-alphabet_length)%26]
					} else {
						rotated_value = alphabet[(index+shiftKey)%26]
					}

					// Add the rotated letter to the string
					// Keep the case of the original letter
					if isUpper {
						rotated_string += string(unicode.ToUpper(rotated_value))
					} else {
						rotated_string += string(rotated_value)
					}
				}
			}
		} else {
			// If the character is not a letter, we just add it to the string
			rotated_string += string(character)
		}
	}

	return rotated_string
}
