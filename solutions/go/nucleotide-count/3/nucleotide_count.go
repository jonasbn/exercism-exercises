package dna

import (
	"errors"
	"strings"
)

// Histogram is a mapping from nucleotide to its count in given DNA.
// Choose a suitable data type.
type Histogram map[rune]int

// DNA holds a strand of DNA consisting of nucleotides
type DNA struct {
	strand string
}

// NewDNA returns DNA struct containing strand
func NewDNA(strand string) *DNA {
	d := DNA{strand: strand}
	return &d
}

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
func (d DNA) Counts() (Histogram, error) {
	h := map[rune]int{'A': 0, 'C': 0, 'G': 0, 'T': 0}

	h['A'] = strings.Count(d.strand, "A")
	h['C'] = strings.Count(d.strand, "C")
	h['G'] = strings.Count(d.strand, "G")
	h['T'] = strings.Count(d.strand, "T")

	if len(d.strand) != sum(h) {
		e := Histogram{}
		return e, errors.New("strand with invalid nucleotides")
	}

	return h, nil
}

func sum(h Histogram) int {
	var sum int

	sum += h['A']
	sum += h['C']
	sum += h['G']
	sum += h['T']

	return sum
}
