package cars

// CalculateWorkingCarsPerHour is used to calculate the ratio of successfully assembled cars per hour
// based on a successrate, takes amount of cars produced and the successrate and return the amount successfully assembled
func CalculateWorkingCarsPerHour(carsPerHour int, successRate float64) float64 {
	return float64(carsPerHour) * (successRate / 100)
}

// CalculateWorkingCarsPerMinute is used to calculate the ratio of successfully assembled cars per minute
// based on a successrate, takes amount of cars produced and the successrate and return the amount successfully assembled
func CalculateWorkingCarsPerMinute(carsPerHour int, successRate float64) int {
	return int(float64(carsPerHour) * (successRate / 100) / 60)
}

// CalculateCost is used to calculate the cost of a production set of cars, it takes an amount of cars and return the cost
func CalculateCost(cars int) uint {
	return uint(((cars / 10) * 95000) + ((cars % 10) * 10000))
}
