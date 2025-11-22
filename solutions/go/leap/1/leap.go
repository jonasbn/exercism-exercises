/*
Package leap can be used to assert leap years for the Gregorian calendar

The package support the proleptic Gregorian calendar, meaning the algoritm handles year prior to the introduction of the Gregorian calendar. In addition the algorithm handles year 0 eventhough it does not exist in the Gregorian calendar, but support for astronomical year numbering is handled.

REF: https://en.wikipedia.org/wiki/Leap_year#Gregorian_calendar
REF: https://en.wikipedia.org/wiki/Astronomical_year_numbering
*/
package leap

// IsLeapYear takes a parameter indicating a year, it returns a boolean
// indicating whether the provided year parameter is a leap year
func IsLeapYear(year int) bool {
	if year == 0 || year%4 > 0 {
		return false
	} else {
		if year%100 == 0 && year%400 > 0 {
			return false
		} else {
			return true
		}
	}
}
