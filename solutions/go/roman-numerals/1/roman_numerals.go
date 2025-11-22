package romannumerals

import (
	"errors"
	"strconv"
	"strings"
)

type romanNumeralSet struct {
	primary   string
	secondary string
	tertiary  string
}

// ToRomanNumeral converts an arabicNumber number to a roman numeral
func ToRomanNumeral(arabicNumber int) (string, error) {

	/* The function does not support arabic numbers below 1, meaning 0 and negative numbers
	   and numbers higher than 3000 */
	if arabicNumber < 1 || arabicNumber > 3000 {
		return "", errors.New("Parameter out of bounds")
	}

	/* We split the arabic number into a splice of strings */
	arabicNumberAsList := reverse(strings.Split(strconv.Itoa(arabicNumber), ""))

	var completeRomanNumeral []string

	for position := 0; position < len(arabicNumberAsList); position++ {
		var romanNumeral []string

		/* We define a small set of roman numerals based on the position in the splice */
		set := resolveRomanNumeralSet(position)

		number, _ := strconv.Atoi(arabicNumberAsList[position])
		remainder := number % 5
		multiplier := number / 5

		/* The roman numeral does not require duplication and we can just append it */
		if multiplier > 0 && remainder < 4 {
			romanNumeral = postfix(romanNumeral, set.secondary)
		}

		if remainder == 4 {
			/* The roman numeral requires subtraction, meaning we append the highest available
			   numeral in the set and prefix the required amount of the lowest available numeral from the set */
			if multiplier > 0 {
				romanNumeral = postfix(romanNumeral, set.tertiary)
				for j := 5 - remainder; j > 0; j-- {
					romanNumeral = prefix(romanNumeral, set.primary)
				}

				/* The roman numeral requires subtraction, meaning we append the highest available
				numeral in the set and prefix the required amount of the lowest available numeral from the set */
			} else {
				romanNumeral = postfix(romanNumeral, set.secondary)
				for k := 5 - remainder; k > 0; k-- {
					romanNumeral = prefix(romanNumeral, set.primary)
				}
			}
		} else if remainder > 0 {
			/* The roman numeral requires addition meaning we append the lowest available numeral from the
			set the required amount of times */
			for l := remainder; l > 0; l-- {
				romanNumeral = postfix(romanNumeral, set.primary)
			}
		}
		/* Numeral assembled and appended to the resulting complete numeral */
		completeRomanNumeral = append(romanNumeral, completeRomanNumeral...)
	}

	return strings.Join(completeRomanNumeral, ""), nil
}

func resolveRomanNumeralSet(i int) romanNumeralSet {
	var r romanNumeralSet

	switch i {
	case 0:
		r = romanNumeralSet{
			primary:   "I",
			secondary: "V",
			tertiary:  "X",
		}
	case 1:
		r = romanNumeralSet{
			primary:   "X",
			secondary: "L",
			tertiary:  "C",
		}
	case 2:
		r = romanNumeralSet{
			primary:   "C",
			secondary: "D",
			tertiary:  "M",
		}
	case 3:
		r = romanNumeralSet{
			primary: "M",
		}
	}

	return r
}

func prefix(romanNumeral []string, s string) []string {
	romanNumeral = append([]string{s}, romanNumeral...)
	return romanNumeral
}

func postfix(romanNumeral []string, s string) []string {
	romanNumeral = append(romanNumeral, s)
	return romanNumeral
}

func reverse(s []string) []string {
	var n []string
	for i := len(s) - 1; i >= 0; i-- {
		n = append(n, s[i])
	}

	return n
}
