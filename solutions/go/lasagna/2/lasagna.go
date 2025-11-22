// Package for calculating lasagna preparation and cooking time
package lasagna

// Holds the estimated time the lasagna has to be in the oven
const OvenTime = 40

// Calculates the remaining time for a lasagna time to be in the oven
func RemainingOvenTime(oventime int) int {
	return OvenTime - oventime
}

// Calculates the time it takes to assemble the lasagna
func PreparationTime(layers int) int {
	return layers * 2
}

// Calculates the time spent to prepare and cook the lasagna
func ElapsedTime(layers int, oventime int) int {
	return PreparationTime(layers) + oventime
}
