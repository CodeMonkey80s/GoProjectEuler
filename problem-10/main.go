package main

import (
	"fmt"
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

const max = 2_000_000

func isPrimeNumber(n int) bool {
	for i := 2; i <= int(math.Floor(float64(n)/2)); i++ {
		if n%i == 0 {
			return false
		}
	}
	return bool(n > 1)
}

func findSumOfPrimesBelowNumber(max int) int {
	var sum int = 0
	var number = 0
	for {
		prime := isPrimeNumber(number)
		if prime {
			sum += number
			fmt.Println(number, "=", sum)
		}
		number++
		if number >= max {
			break
		}
	}
	return sum
}

func main() {
	v := findSumOfPrimesBelowNumber(max)
	fmt.Println(v)
}
