// Package allowss for booking and scheduling appointment in a beauty salon
package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {

	// Output example: 7/25/2019 13:45:00
	const longForm string = "1/2/2006 15:04:05"
	t, err := time.Parse(longForm, date)
	if err != nil {
		fmt.Printf("Invalid date specified: %s", date)
	}

	// Output example: 2019-07-25 13:45:00 +0000 UTC
	return t
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {

	// Input example: December 9, 2112 11:59:59
	const longForm string = "January 2, 2006 15:04:05"
	t, err := time.Parse(longForm, date)
	if err != nil {
		fmt.Printf("Invalid date specified: %s", date)
	}

	if t.Before(time.Now()) {
		return true
	}

	return false
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {

	// Input example: Thursday, May 13, 2010 20:32:00
	const longForm string = "Monday, January 2, 2006 15:04:05"
	t, err := time.Parse(longForm, date)
	if err != nil {
		fmt.Printf("Invalid date specified: %s", date)
	}

	if t.Hour() >= 12 && t.Hour() < 18 {
		return true
	}

	return false
}

// Description returns a formatted string of the appointment time
func Description(date string) string {

	// Output example: 7/25/2019 13:45:00
	const longForm string = "1/2/2006 15:04:05"
	t, err := time.Parse(longForm, date)
	if err != nil {
		fmt.Printf("Invalid date specified: %s", date)
	}

	description := fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %d:%d.", t.Weekday(), t.Month(), t.Day(), t.Year(), t.Hour(), t.Minute())

	// "You have an appointment on Thursday, July 25, 2019, at 13:45."
	return description
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {

	anniversaryDate := time.Date(
		2025, 9, 15, 0, 0, 0, 0, time.UTC)

	// Output: 2020-09-15 00:00:00 +0000 UTC
	return anniversaryDate
}
