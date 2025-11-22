// Package twofer implements a simple library communicating a sharing strategy for two parties
package twofer

// ShareWith returns a string communicating either anonymous or named sharing
// The optional parameter can be a name of the party to share with or nothing
// for a more anonymous communication
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}

	return "One for " + name + ", one for me."
}
