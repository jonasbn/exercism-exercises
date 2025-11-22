package wordcount

import (
	"regexp"
	"strings"
)

func WordCount(phrase string) map[string]int {
	counts := make(map[string]int)

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
