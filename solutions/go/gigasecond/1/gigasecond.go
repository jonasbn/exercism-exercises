// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package gigasecond offers a function to add gigaseconds
package gigasecond

// import path for the time package from the standard library
import (
	"time"
)

// AddGigasecond adds a gigasecond to a given date and returns the extended date
func AddGigasecond(t time.Time) time.Time {

	futuredate := t.Add(time.Second * 1000000000)

	return futuredate
}
