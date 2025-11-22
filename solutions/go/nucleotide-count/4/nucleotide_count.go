package dna

import "errors"

type Histogram map[rune]int

type DNA string

func (d DNA) Counts() (Histogram, error) {
	var h Histogram = map[rune]int{'A': 0, 'C': 0, 'G': 0, 'T': 0}
	for _, l := range d {

		switch l {
		case 'A', 'C', 'G', 'T':
			h[l]++
		default:
			return nil, errors.New("INVALID")
		}
	}
	return h, nil
}
