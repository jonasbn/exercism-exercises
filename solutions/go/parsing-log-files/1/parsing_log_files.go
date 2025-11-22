package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)

	if re.MatchString(text) {
		return true
	}

	return false
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[[:^alpha:]]*?>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	var count int

	re := regexp.MustCompile(`(?i)".*password.*"`)

	for _, line := range lines {
		if re.MatchString(line) {
			count++
		}
	}

	return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end\-of\-line\d+`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {

	var parsedLines []string

	reFind := regexp.MustCompile(`User\s+(\w+)`)
	reReplace := regexp.MustCompile(`^`)

	for _, line := range lines {
		s := reFind.FindStringSubmatch(line)
		if s != nil {
			replacement := fmt.Sprintf("[USR] %s ", s[1])
			line = reReplace.ReplaceAllString(line, replacement)
		}

		parsedLines = append(parsedLines, line)
	}

	return parsedLines
}
