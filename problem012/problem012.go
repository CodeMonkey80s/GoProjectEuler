package problem012

import (
	"math"
)

const maxDivisors int = 500

func countFactors(n int) int {
	var count = 0
	var sqrt = int(math.Sqrt(float64(n)))
	for i := 1; i <= sqrt; i++ {
		if n%i == 0 {
			count += 2
		}
		if sqrt*sqrt == n {
			count--
		}
	}
	return count
}

func findHighlyDivisibleTriangularNumber() int {
	var n = 0
	var i = 0
	var factors = 0
	for {
		factors = countFactors(n)
		if factors > maxDivisors {
			break
		}
		n += i
		i++
	}
	return n
}

func Solve() int {
	return findHighlyDivisibleTriangularNumber()
}
