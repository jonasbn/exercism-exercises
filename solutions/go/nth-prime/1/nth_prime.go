package prime

import "fmt"

// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
	if n < 1 {
		err := fmt.Errorf("n should be a positive integer")
		return 0, err
	}

	var primes []int

	for i := 0; len(primes) < n; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}

	return primes[n-1], nil
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}
