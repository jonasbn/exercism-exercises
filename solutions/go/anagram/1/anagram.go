package anagram

import (
	"strings"
)

// Detect compares a fixture string against a list of strings for
// detection of anagrams, returns a list of valid anagrams if any
func Detect(a string, b []string) []string {
	var detected []string

	var replacement string
	for i := 0; i < len(b); i++ {
		word := strings.ToLower(b[i])

		// Our words are not same length
		if len(a) != len(word) {
			continue
		}

		// Our words already match (case insensitive)
		if strings.ToLower(a) == word {
			continue
		}

		// Eliminate characters for anagram detection
		for j := 0; j < len(word); j++ {
			letter := strings.ToLower(string(a[j]))
			replacement = strings.Replace(word, letter, "_", 1)
			word = replacement

		}
		// Success, we have an anagram, we keep the original
		if strings.Count(word, "_") == len(b[i]) {
			detected = append(detected, b[i])
		}
	}

	return detected
}
