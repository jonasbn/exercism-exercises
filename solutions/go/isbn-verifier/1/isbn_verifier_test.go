package isbn

import (
	"fmt"
	"testing"
)

func TestIsValidISBN(t *testing.T) {
	for _, test := range testCases {
		t.Run(test.description, func(t *testing.T) {
			observed := IsValidISBN(test.isbn)
			if observed != test.expected {
				t.Errorf("IsValidISBN(%q)\nExpected: %t, Actual: %t",
					test.isbn, test.expected, observed)
			}
		})
	}
}

func TestConvertISBN10to13(t *testing.T) {

	var testConversions = []struct {
		isbn        string
		expected    string
		err         error
		description string
	}{
		{"3-598-21508-8", "9783598215087", nil, "succesfully converted isbn 10 number"},
		{"0-306-40615-2", "9780306406157", nil, "succesfully converted isbn 10 number"},
		{"3-598-21507-X", "9783598215070", nil, "succesfully converted isbn 10 number"},
		{"3598215088", "9783598215087", nil, "succesfully converted isbn 10 number"},
		{"", "", fmt.Errorf("specified ISBN: %s insufficient for conversion", ""), "empty parameter should result in error"},
		{"00", "", fmt.Errorf("specified ISBN: %s insufficient for conversion", "00"), "too short parameter should result in error"},
		{"134456729", "", fmt.Errorf("specified ISBN: %s insufficient for conversion", "134456729"), "input is 9 characters"},
		{"3132P34035", "", fmt.Errorf("specified ISBN: %s insufficient for conversion", "3132P34035"), "invalid characters are not ignored"},
		{"98245726788", "", fmt.Errorf("specified ISBN: %s insufficient for conversion", "98245726788"), "input is too long but contains a valid isbn"},
	}

	for _, test := range testConversions {
		t.Run(test.description, func(t *testing.T) {
			observed, err := ConvertISBN10to13(test.isbn)
			if observed != test.expected || (err != nil && err.Error() != test.err.Error()) {
				t.Errorf("ConvertISBN10to13(%q)\nExpected: >%s< and >%s<, Actual: >%s< and >%s<",
					test.isbn, test.expected, test.err, observed, err)

			}
		})
	}
}

func TestCalculateISBN13checksum(t *testing.T) {

	var testCalculations = []struct {
		isbn        string
		expected    string
		err         error
		description string
	}{
		{"978-0-306-40615-7", "7", nil, "succesfully calculated checksum for isbn 13 number"},
		{"978-3-16-148410-0", "0", nil, "succesfully calculated checksum for isbn 13 number"},
		{"9783598215087", "7", nil, "succesfully calculated checksum for isbn 13 number"},
		{"9783598215070", "0", nil, "succesfully calculated checksum for isbn 13 number"},
		{"", "", fmt.Errorf("specified ISBN: %s insufficient for checksum calculation", ""), "empty parameter should result in error"},
		{"00", "", fmt.Errorf("specified ISBN: %s insufficient for checksum calculation", "00"), "too short parameter should result in error"},
		{"134456729", "", fmt.Errorf("specified ISBN: %s insufficient for checksum calculation", "134456729"), "input is 9 characters"},
		{"3-598-21507-X", "", fmt.Errorf("specified ISBN: %s insufficient for checksum calculation", "3-598-21507-X"), "invalid characters are not ignored"},
		{"97835982150700", "", fmt.Errorf("specified ISBN: %s insufficient for checksum calculation", "97835982150700"), "input is too long but contains a valid isbn"},
	}

	for _, test := range testCalculations {
		t.Run(test.description, func(t *testing.T) {
			observed, err := CalculateISBN13checksum(test.isbn)
			if observed != test.expected || (err != nil && err.Error() != test.err.Error()) {
				t.Errorf("CalculateISBN13checksum(%q)\nExpected: >%s< and >%s<, Actual: >%s< and >%s<",
					test.isbn, test.expected, test.err, observed, err)
			}
		})
	}
}

func TestCalculateISBN10checksum(t *testing.T) {

	var testCalculations = []struct {
		isbn        string
		expected    string
		err         error
		description string
	}{
		{"3-598-21508-8", "8", nil, "succesfully calculated checksum for isbn 10 number"},
		{"0-306-40615-2", "2", nil, "succesfully calculated checksum for isbn 10 number"},
		{"3-598-21507-X", "X", nil, "succesfully calculated checksum for isbn 10 number"},
		{"3598215088", "8", nil, "succesfully calculated checksum for isbn 10 number"},
		{"", "", fmt.Errorf("specified ISBN: %s insufficient for checksum calculation", ""), "empty parameter should result in error"},
		{"00", "", fmt.Errorf("specified ISBN: %s insufficient for checksum calculation", "00"), "too short parameter should result in error"},
		{"134456729", "", fmt.Errorf("specified ISBN: %s insufficient for conversion", "134456729"), "input is 9 characters"},
		{"3132P34035", "", fmt.Errorf("specified ISBN: %s insufficient for conversion", "3132P34035"), "invalid characters are not ignored"},
		{"98245726788", "", fmt.Errorf("specified ISBN: %s insufficient for conversion", "98245726788"), "input is too long but contains a valid isbn"},
	}

	for _, test := range testCalculations {
		t.Run(test.description, func(t *testing.T) {
			observed, err := CalculateISBN10checksum(test.isbn)
			if observed != test.expected {
				t.Errorf("CalculateISBN10checksum(%q)\nExpected: >%s< and >%s<, Actual: >%s< and >%s<",
					test.isbn, test.expected, test.err, observed, err)
			}
		})
	}
}

func BenchmarkIsValidISBN(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, n := range testCases {
			IsValidISBN(n.isbn)
		}
	}
}
