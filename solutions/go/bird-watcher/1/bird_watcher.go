package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {

	var totalDailyBirdCount int = 0

	for _, dailyBirdCount := range birdsPerDay {
		totalDailyBirdCount += dailyBirdCount
	}

	return totalDailyBirdCount
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {

	const weekLength int = 7

	weekStart := weekLength*week - weekLength
	weekEnd := weekStart + weekLength + 1

	if len(birdsPerDay) < weekEnd {
		weekEnd = len(birdsPerDay)
	}

	birdsPerDayByWeek := birdsPerDay[weekStart:weekEnd]

	return TotalBirdCount(birdsPerDayByWeek)
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {

	for i := 0; i < len(birdsPerDay); i += 2 {
		birdsPerDay[i]++
	}

	return birdsPerDay
}
