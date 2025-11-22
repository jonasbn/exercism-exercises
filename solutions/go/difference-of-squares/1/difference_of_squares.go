package diffsquares

// REF: https://www.mathblog.dk/project-euler-problem-6/

func SumOfSquares(n int) int {
	return (n * (n + 1) * (2*n + 1)) / 6
}

func SquareOfSum(n int) int {
	return (n * (n + 1) / 2) * (n * (n + 1) / 2)
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
