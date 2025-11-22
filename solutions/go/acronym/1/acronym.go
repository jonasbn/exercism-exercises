// Abbreviate should have a comment documenting it.
package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(input string) string {
	var output []string

	words := regexp.MustCompile("[[:space:]]+|[-_]+").Split(input, -1)

	for _, word := range words {
		letter := regexp.MustCompile("^([[:alpha:]]{1})").FindAllString(word, -1)

		output = append(output, strings.ToUpper(strings.Join(letter, "")))
	}

	return strings.Join(output, "")
}
