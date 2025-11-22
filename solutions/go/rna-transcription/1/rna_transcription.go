// Package strand offers functionality to transcribe DNA strands to the complement RNA strand
package strand

import (
	"strings"
)

// ToRNA transcribes DNA to RNA
func ToRNA(dna string) string {

	transcribe := func(r rune) rune {
		switch r {
		case 'G':
			return 'C'
		case 'C':
			return 'G'
		case 'T':
			return 'A'
		case 'A':
			return 'U'
		}
		return r
	}

	return strings.Map(transcribe, dna)
}
