// Package diffsquares offers functions to calculate: sum of squares, square of a sum and the difference between the two
// This is an implementation of the problem 6 from Project Euler
// See: https://projecteuler.net/problem=6
// The implementation is based on a solution my Ken Ward
// See: https://www.mathblog.dk/project-euler-problem-6/
package diffsquares

// SumOfSquares calculates the sum of squares for natural numbers
func SumOfSquares(n int) int {
	return (n * (n + 1) * (2*n + 1)) / 6
}

// SquareOfSum calculates the the square of a sum for natural numbers
func SquareOfSum(n int) int {
	return (n * (n + 1) / 2) * (n * (n + 1) / 2)
}

// Difference calculates the difference between the SquareOfSum and SumOfSquares
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
