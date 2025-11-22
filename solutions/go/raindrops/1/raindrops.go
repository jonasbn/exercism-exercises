// Package raindrops converts integers into strings resembling raindrop sounds
// based in a set of calculation rules
package raindrops

import (
	"strconv"
	"strings"
)

// Convert outputs raindrops string based on integer input
func Convert(n int) string {

	var s []string

	// The order of the ifs is significant
	if n%3 == 0 {
		s = append(s, "Pling")
	}

	if n%5 == 0 {
		s = append(s, "Plang")
	}

	if n%7 == 0 {
		s = append(s, "Plong")
	}

	// We have no raindrops, so we return the original integer
	// We have to convert it to string
	if len(s) == 0 {
		s = append(s, strconv.Itoa(n))
	}

	// Convert our slice to a string
	return strings.Join(s, "")
}
