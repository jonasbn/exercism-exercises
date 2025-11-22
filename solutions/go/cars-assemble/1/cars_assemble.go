package cars

// SuccessRate is used to calculate the ratio of an item being created without
// error for a given speed
func SuccessRate(speed int) float64 {

	var successRate float64

	switch {
	case speed > 8:
		successRate = 0.77
	case speed > 4:
		successRate = 0.9
	case speed > 0:
		successRate = 1.0
	default:
		successRate = 0.0
	}

	return successRate
}

// CalculateProductionRatePerHour for the assembly line, taking into account
// its success rate
func CalculateProductionRatePerHour(speed int) float64 {
	return float64(speed*221) * SuccessRate(speed)
}

// CalculateProductionRatePerMinute describes how many working items are
// produced by the assembly line every minute
func CalculateProductionRatePerMinute(speed int) int {
	return int(CalculateProductionRatePerHour(speed) / 60)
}

// CalculateLimitedProductionRatePerHour describes how many working items are
// produced per hour with an upper limit on how many can be produced per hour
func CalculateLimitedProductionRatePerHour(speed int, limit float64) float64 {
	productionRatePerHour := CalculateProductionRatePerHour(speed)

	if productionRatePerHour > limit {
		return limit
	}

	return productionRatePerHour
}
