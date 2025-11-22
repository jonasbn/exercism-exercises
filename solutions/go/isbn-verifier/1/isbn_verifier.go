package isbn

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func cleanISBNString(isbn string) []string {
	re := regexp.MustCompile(`[^0-9X]`)
	cleanStr := re.ReplaceAllString(isbn, "")

	s := strings.SplitAfter(cleanStr, "")

	return s
}

func IsValidISBN(isbn string) bool {

	s := cleanISBNString(isbn)

	switch len(s) {
	case 10:
		return isValidISBN10(s)
	case 13:
		return isValidISBN13(s)
	default:
		return false
	}
}

func CalculateISBN13checksum(isbn string) (string, error) {

	s := cleanISBNString(isbn)

	if len(s) != 13 {
		return "", fmt.Errorf("specified ISBN: %s insufficient for checksum calculation", isbn)
	}

	sum := 0

	for i := 0; i < 12; i++ {
		n, err := strconv.Atoi(s[i])
		if err != nil {
			return "-1", err
		}

		if (i+1)%2 == 0 {
			sum += n * 3
		} else {
			sum += n
		}
	}

	r := sum % 10

	c := 10 - r
	checksum := c % 10

	return strconv.Itoa(checksum), nil
}

func CalculateISBN10checksum(isbn string) (string, error) {

	s := cleanISBNString(isbn)

	if len(s) != 10 {
		return "", fmt.Errorf("specified ISBN: %s insufficient for checksum calculation", isbn)
	}

	if isValidISBN10(s) {
		sum := 0

		for i := 0; i < 9; i++ {
			n, err := strconv.Atoi(s[i])
			if err != nil {
				if s[i] == "X" && i == 8 {
					n = 10
				} else {
					return "", err
				}
			}
			sum += n * (i + 1)
		}

		c := sum % 11
		var checksum string

		if c == 10 {
			checksum = "X"
		} else {
			checksum = strconv.Itoa(c)
		}

		return checksum, nil
	} else {
		return "", fmt.Errorf("unable to calculate checksum for ISBN: %s", isbn)
	}
}

func ConvertISBN10to13(isbn string) (string, error) {

	s := cleanISBNString(isbn)

	if len(s) != 10 {
		return "", fmt.Errorf("specified ISBN: %s insufficient for conversion", isbn)
	}

	if isValidISBN10(s) {

		p := []string{"9", "7", "8"}
		p = append(p, s...)

		c, err := CalculateISBN13checksum(strings.Join(p, ""))
		if err != nil {
			return "", err
		}

		p[12] = c

		if isValidISBN13(p) {
			return strings.Join(p, ""), nil
		} else {
			return "", fmt.Errorf("generated ISBN not valid %s", strings.Join(p, ""))
		}
	} else {
		return "", fmt.Errorf("unable to convert ISBN: %s", isbn)
	}
}

func isValidISBN10(isbn []string) bool {

	if len(isbn) != 10 {
		return false
	}

	sum := 0
	for i := 0; i < 10; i++ {
		n, err := strconv.Atoi(isbn[i])
		if err != nil {
			if isbn[i] == "X" && i == 9 {
				n = 10
			} else {
				return false
			}
		}
		sum += n * (i + 1)
	}

	return sum%11 == 0
}

func isValidISBN13(isbn []string) bool {

	if len(isbn) != 13 {
		return false
	}

	sum := 0
	for i := 0; i < 13; i++ {
		n, err := strconv.Atoi(isbn[i])
		if err != nil {
			return false
		}

		if n == 0 {
			n = 10
		}

		if (i+1)%2 == 0 {
			sum += n * 3
		} else {
			sum += n
		}
	}

	return sum%10 == 0
}
