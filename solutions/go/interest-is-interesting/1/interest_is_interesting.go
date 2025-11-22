package interest

const NegativeInterestRate float32 = 3.213
const LowPositiveInterestRate float32 = 0.5
const MiddlePositiveInterestRate float32 = 1.621
const HighPositiveInterestRate float32 = 2.475

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	var interestRate float32

	switch {
	case balance < 0:
		interestRate = NegativeInterestRate
	case balance >= 0 && balance < 1000:
		interestRate = LowPositiveInterestRate
	case balance >= 1000 && balance < 5000:
		interestRate = MiddlePositiveInterestRate
	case balance >= 5000:
		interestRate = HighPositiveInterestRate
	}

	return interestRate
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	return balance * float64(InterestRate(balance)/100)
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return balance + Interest(balance)
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance:
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	var years int = 0

	for {
		if balance >= targetBalance {
			break
		}
		balance = AnnualBalanceUpdate(balance)
		years++
	}
	return years
}
