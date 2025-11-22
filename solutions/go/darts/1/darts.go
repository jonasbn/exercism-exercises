package darts

import "math"

func Score(x float64, y float64) int {

	distance := calculate_distance(x, y)

	switch {
	case distance <= 1:
		return 10
	case distance <= 5:
		return 5
	case distance <= 10:
		return 1
	default:
		return 0
	}
}

func calculate_distance(x float64, y float64) float64 {
	// REF: https://en.wikipedia.org/wiki/Distance
	// REF: https://byjus.com/maths/distance-between-points/
	return math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
}
