package problem006

import (
	"math"
)

// Problem 6
//
// URL: https://projecteuler.net/problem=6
//
// ============================================================================

const maxNumbers int = 100

func sumOfSquares(max int) int {
	var sum = 0
	for i := 1; i <= max; i++ {
		sum += i * i
	}
	return sum
}

func squareOfSums(max int) int {
	var sum = 0
	for i := 1; i <= max; i++ {
		sum += i
	}
	return int(math.Pow(float64(sum), 2))
}

func Solve() int {
	x := sumOfSquares(maxNumbers)
	y := squareOfSums(maxNumbers)
	return y - x
}
