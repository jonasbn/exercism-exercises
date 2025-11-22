package luhn

import (
	"strconv"
	"strings"
	"unicode"
)

func Valid(id string) bool {

	// we remove spaces
	letters := strings.FieldsFunc(id, func(r rune) bool {
		return unicode.IsSpace(r)
	})

	// we convert to a string again
	s := strings.Join(letters, "")

	// we do not handle strings with the length of 1 and below
	if len(s) <= 1 {
		return false
	}

	// we split into a string slice
	letters = strings.Split(s, "")

	var sum int
	// we have to handle both even and odd lenght strings
	// REF: https://rosettacode.org/wiki/Luhn_test_of_credit_card_numbers#Go
	odd := len(letters) & 1

	// we traverse the slice from the back to the front
	for i := len(letters) - 1; i >= 0; i-- {
		number, err := strconv.Atoi(string(letters[i]))
		if err != nil {
			return false
		}
		if i&1 == odd {
			number = number * 2
			if number > 9 {
				number = number - 9
			}
		}
		sum += number
	}

	if sum%10 == 0 {
		return true
	}

	return false
}
