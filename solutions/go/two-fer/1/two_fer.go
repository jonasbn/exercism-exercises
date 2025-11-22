/*
Package twofer implements a simple library communicating a sharing strategy for two parties
*/
package twofer

import "fmt"

// The optional parameter can be a name of the party to share with or nothing
// for a more anonymous communication
//
// ShareWith returns a string communicating either anonymous or named sharing
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}

	return fmt.Sprintf("One for %s, one for me.", name)
}
