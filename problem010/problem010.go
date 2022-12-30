package problem010

import (
	"math"
)

// Problem 10
//
// URL: https://projecteuler.net/problem=10
//
// The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
//
// Find the sum of all the primes below two million.
//
// ============================================================================

// const max = 2_000_000
const max = 1000

func isPrimeNumber(n int) bool {
	for i := 2; i <= int(math.Floor(float64(n)/2)); i++ {
		if n%i == 0 {
			return false
		}
	}
	return n > 1
}

func findSumOfPrimesBelowNumber(max int) int {
	var sum = 0
	var number = 0
	for {
		prime := isPrimeNumber(number)
		if prime {
			sum += number
		}
		number++
		if number >= max {
			break
		}
	}
	return sum
}

func Solve() int {
	return findSumOfPrimesBelowNumber(max)
}
