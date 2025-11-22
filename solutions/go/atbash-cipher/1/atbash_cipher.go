package atbash

import (
	"regexp"
	"strings"
)

var alphabet = []string{
	"a", "b", "c", "d", "e", "f", "g", "h", "i", "j",
	"k", "l", "m", "n", "o", "p", "q", "r", "s", "t",
	"u", "v", "w", "x", "y", "z",
}

var reversedAlphabet = []string{
	"z", "y", "x", "w", "v", "u", "t", "s", "r", "q",
	"p", "o", "n", "m", "l", "k", "j", "i", "h", "g",
	"f", "e", "d", "c", "b", "a",
}

func Atbash(s string) string {

	s = strings.ToLower(s)

	re := regexp.MustCompile(`[^a-z0-9]*`)
	s = re.ReplaceAllString(s, "")

	m := make(map[string]string)
	for i, letter := range alphabet {
		m[letter] = reversedAlphabet[i]
	}
	var r []string
	counter := 1
	for _, letter := range strings.Split(s, "") {
		if val, ok := m[letter]; ok {
			r = append(r, val)
		} else {
			r = append(r, letter)
		}
		if counter%5 == 0 && counter != 0 {
			r = append(r, " ")
		}
		counter++
	}

	result := strings.Join(r, "")
	result = strings.TrimSpace(result)

	return result
}
