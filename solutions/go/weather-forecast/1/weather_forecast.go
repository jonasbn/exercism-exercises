// Package weather implments functionality to do weather forecasts.
package weather

// CurrentCondition specifies a string representing the current weather condition.
var CurrentCondition string

// CurrentLocation specifies a string representing the current location.
var CurrentLocation string

// Forecast returns a string with the current forecast, based on the two input parameters: city and condition, both strings.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
