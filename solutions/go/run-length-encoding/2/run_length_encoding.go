package encode

import (
	"strconv"
	"strings"
)

func RunLengthEncode(input string) string {
	var output []string

	if len(input) > 0 {
		count := 0
		previous := ""

		for _, current := range strings.Split(input, "") {
			if current != previous && previous != "" {
				if count > 1 {
					output = append(output, strconv.Itoa(count))
				}
				output = append(output, previous)
				count = 0
				previous = ""
			}
			count = count + 1
			previous = current
		}
		if count > 1 {
			output = append(output, strconv.Itoa(count))
		}
		output = append(output, previous)
	}

	return strings.Join(output, "")
}

func RunLengthDecode(input string) string {
	var output []string
	var number []string
	previous := ""

	for _, current := range strings.Split(input, "") {

		if isDigit(current) {
			number = append(number, current)
		} else {

			if isDigit(previous) {
				multiplier, _ := strconv.Atoi(strings.Join(number, ""))
				output = append(output, strings.Repeat(current, multiplier))
				number = nil
			} else {
				output = append(output, current)
			}
		}
		previous = current
	}

	return strings.Join(output, "")
}

func isDigit(s string) (is_digit bool) {

	_, e := strconv.Atoi(s)

	return e == nil
}
