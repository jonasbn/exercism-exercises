package protein

import (
	"errors"
)

// ErrInvalidBase blah
var ErrInvalidBase = errors.New("InvalidBase")

// ErrStop blah
var ErrStop = errors.New("STOP")

// FromCodon converts a codon to it's human readable protein name
func FromCodon(codon string) (string, error) {
	protein := ""

	switch codon {
	case "AUG":
		protein = "Methionine"
		break
	case "UUU":
		protein = "Phenylalanine"
		break
	case "UUC":
		protein = "Phenylalanine"
		break
	case "UUA":
		protein = "Leucine"
		break
	case "UUG":
		protein = "Leucine"
		break
	case "UCU":
		protein = "Serine"
		break
	case "UCC":
		protein = "Serine"
		break
	case "UCA":
		protein = "Serine"
		break
	case "UCG":
		protein = "Serine"
		break
	case "UAU":
		protein = "Tyrosine"
		break
	case "UAC":
		protein = "Tyrosine"
		break
	case "UGU":
		protein = "Cysteine"
		break
	case "UGC":
		protein = "Cysteine"
		break
	case "UGG":
		protein = "Tryptophan"
		break
	case "UAA":
		return "", ErrStop
	case "UAG":
		return "", ErrStop
	case "UGA":
		return "", ErrStop
	default:
		return "", ErrInvalidBase
	}

	return protein, nil
}

// FromRNA converts a RNA to it's human readable protein name by translating it's codons
func FromRNA(rna string) ([]string, error) {
	var proteins []string

	if len(rna)%3 != 0 {
		return proteins, ErrInvalidBase
	}

	for i := 0; i < len(rna); i = i + 3 {
		protein, err := FromCodon(rna[i : i+3])

		if err == nil {
			proteins = append(proteins, protein)
		} else if err == ErrStop {
			return proteins, nil
		} else {
			return proteins, err
		}
	}
	return proteins, nil
}
