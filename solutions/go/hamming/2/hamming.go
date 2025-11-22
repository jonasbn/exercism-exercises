// Package hamming offers function to calculate Hamming distance between two DNA strands
package hamming

import (
	"errors"
	"unicode/utf8"
)

// Distance calculate string distrance between to strings or errors
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("lengths of provided strings differs")
	}

	var distance int
	for _, runeFromA := range a {
		runeFromB, runeFromBSize := utf8.DecodeRuneInString(b)
		b = b[runeFromBSize:]

		if runeFromB != runeFromA {
			distance++
		}
	}
	return distance, nil
}
