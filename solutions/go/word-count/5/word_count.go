package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	counts := Frequency{}

	re_split := regexp.MustCompile(`[,.:!@$%^&\s]+`)
	s := re_split.Split(phrase, -1)

	for _, v := range s {
		re_replace := regexp.MustCompile(`^'|'$`)
		v = re_replace.ReplaceAllString(v, "")
		if v != "" {
			counts[strings.ToLower(v)]++
		}
	}

	return counts
}
