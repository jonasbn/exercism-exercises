// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import "regexp"

// Hey should have a comment documenting it.
func Hey(remark string) string {

	re_lower := regexp.MustCompile(`[[:lower:]]+`)
	re_upper := regexp.MustCompile(`[[:upper:]]+`)
	re_question := regexp.MustCompile(`\?\s*$`)
	re_empty := regexp.MustCompile(`^$`)
	re_whitespace := regexp.MustCompile(`^[[:space:]]+$`)

	if re_empty.MatchString(remark) || re_whitespace.MatchString(remark) {
		return "Fine. Be that way!"
	} else if re_question.MatchString(remark) && re_upper.MatchString(remark) && !re_lower.MatchString(remark) {
		return "Calm down, I know what I'm doing!"
	} else if re_question.MatchString(remark) {
		return "Sure."
	} else if re_upper.MatchString(remark) && !re_lower.MatchString(remark) {
		return "Whoa, chill out!"
	} else {
		return "Whatever."
	}
}
