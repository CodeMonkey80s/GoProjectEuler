package problem003

import (
	"math"
)

// Problem 3
//
// URL: https://projecteuler.net/problem=3
//
// The prime factors of 13195 are 5, 7, 13 and 29.
//
// What is the largest prime factor of the number y ?
//
// ============================================================================

const number int = 600851475143

func primeFactors(n int) int {
	var largest = 0
	for n%2 == 0 {
		n /= 2
	}
	max := int(math.Sqrt(float64(n)))
	for i := 3; i <= max; i += 2 {
		for n%i == 0 {
			if i > largest {
				largest = i
			}
			n /= i
		}
	}
	if n > 2 {
		if n > largest {
			largest = n
		}
	}
	return largest
}

func Solve() int {
	return primeFactors(number)
}
