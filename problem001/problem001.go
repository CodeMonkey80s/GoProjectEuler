package problem001

// Problem 1
//
// URL: https://projecteuler.net/problem=1
//
// If we list all the natural numbers below 10 that are multiples of 3 or 5,
// we get 3, 5, 6 and 9. The sum of these multiples is 23.
//
// Find the sum of all the multiples of 3 or 5 below 1000.
//
// ============================================================================

const maxNumber int = 1000

func calculate(max int) int {
	var sum = 0
	var i = 0
	for i = 1; i < max; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	return sum
}

func Solve() int {
	return calculate(maxNumber)
}
