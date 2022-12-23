package main

import (
	"fmt"
	"math"
)

// Problem 6
//
// URL: https://projecteuler.net/problem=6
//
// ============================================================================

const maxNumbers int = 100

func sumOfSquares(max int) int {
	var sum int = 0
	for i := 1; i <= max; i++ {
		sum += (i * i)
	}
	return sum
}

func squareOfSums(max int) int {
	var sum int = 0
	for i := 1; i <= max; i++ {
		sum += i
	}
	return int(math.Pow(float64(sum), 2))
}

func main() {
	x := sumOfSquares(maxNumbers)
	y := squareOfSums(maxNumbers)
	fmt.Println("Difference", y-x)
}
