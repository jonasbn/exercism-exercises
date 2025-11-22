// Package accumulate offers a poor man's map function
package accumulate

// https://tour.golang.org/moretypes/25
type converter func(string) string

// Accumulate converts a collection using the provided converter and returns the convertet collection
func Accumulate(c []string, fn converter) []string {

	var s []string

	for i := 0; i < len(c); i++ {
		s = append(s, fn(c[i]))
	}
	return s
}
