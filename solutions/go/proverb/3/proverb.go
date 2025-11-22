// Package proverb generates proverbs based on a list of requested rhymes
package proverb

import "fmt"

// Proverb prints a proverb based on the contents of the rhymes requested
func Proverb(rhyme []string) []string {

	// Default to original proverb
	var word = "nail"

	// Traditional proverb
	var traditionalVerses = map[string]string{
		"nail":    "",
		"shoe":    "nail",
		"horse":   "shoe",
		"rider":   "horse",
		"message": "rider",
		"battle":  "message",
		"kingdom": "battle",
	}

	// Modern proverb, expecting pin as new identification
	var modernVerses = map[string]string{
		"pin":     "",
		"gun":     "pin",
		"soldier": "gun",
		"battle":  "soldier",
	}

	var verses = map[string]string{}

	if contains(word, rhyme) {
		verses = traditionalVerses
	} else {
		// Setting pin as unique identification of last verse
		word = "pin"
		verses = modernVerses
	}

	var proverb []string

	for _, v := range rhyme {
		// We skip the verse associated with our proverb identication
		// for handling order properly, see below for last append operation
		// of the final verse
		if v != word {
			proverb = append(proverb, fmt.Sprintf("For want of a %s the %s was lost.", verses[v], v))
		}
	}

	// We append the final verse
	if len(rhyme) > 0 {
		proverb = append(proverb, fmt.Sprintf("And all for the want of a %s.", word))
	}

	return proverb
}

// contains check a string slice of words for the presence of a given word
func contains(word string, words []string) bool {
	for i := 0; i < len(words); i++ {
		if words[i] == word {
			return true
		}
	}

	return false
}
