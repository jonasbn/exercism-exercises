// Package proverb generates proverbs based on a list of requested rhymes
package proverb

// Proverb prints a proverb based on the contents of the rhymes requested
func Proverb(rhyme []string) []string {

	// Default to original proverb
	var word string = "nail"

	// Original proverb
	var ancientVerses = map[string]string{
		"nail":    "And all for the want of a nail.",
		"shoe":    "For want of a nail the shoe was lost.",
		"horse":   "For want of a shoe the horse was lost.",
		"rider":   "For want of a horse the rider was lost.",
		"message": "For want of a rider the message was lost.",
		"battle":  "For want of a message the battle was lost.",
		"kingdom": "For want of a battle the kingdom was lost.",
	}

	// Modern proverb, expecting pin as new identification
	var modernVerses = map[string]string{
		"pin":     "And all for the want of a pin.",
		"gun":     "For want of a pin the gun was lost.",
		"soldier": "For want of a gun the soldier was lost.",
		"battle":  "For want of a soldier the battle was lost.",
	}

	var verses = map[string]string{}

	if contains(word, rhyme) {
		verses = ancientVerses
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
			proverb = append(proverb, verses[v])
		}
	}

	// We append the last verse
	if len(rhyme) > 0 {
		proverb = append(proverb, verses[word])
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
