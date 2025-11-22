// Package hamming offers function to calculate Hamming distance between two DNA strands
package hamming

import (
	"errors"
	"strings"
)

// Distance calculate string distrance between to strings or errors
func Distance(a, b string) (int, error) {
	var distance int = 0

	if len(a) != len(b) {
		return 0, errors.New("lengths of provided strings differs")
	}
	s1 := strings.Split(a, "")
	s2 := strings.Split(b, "")
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			distance++
		}
	}
	return distance, nil
}
