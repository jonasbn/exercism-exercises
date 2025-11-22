package logs

import (
	"regexp"
	"unicode/utf8"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	var application string

	//re := regexp.MustCompile(`(U+2757|U+1F50D|U+2600)`)
	re := regexp.MustCompile(`(❗|🔍|☀)`)
	c := re.FindString(log)

	switch c {
	case "❗":
		application = "recommendation"
	case "🔍":
		application = "search"
	case "☀":
		application = "weather"
	default:
		application = "default"
	}

	return application
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	var b []byte

	for len(log) > 0 {
		r, size := utf8.DecodeRuneInString(log)

		if r == oldRune {
			newSize := utf8.RuneLen(newRune)
			buf := make([]byte, newSize)
			utf8.EncodeRune(buf, newRune)
			b = append(b, buf...)
		} else {
			buf := make([]byte, size)
			utf8.EncodeRune(buf, r)
			b = append(b, buf...)
		}

		log = log[size:]
	}

	return string(b)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
