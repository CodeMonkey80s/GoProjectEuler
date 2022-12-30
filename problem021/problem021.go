package problem021

// Problem 21
//
// URL: https://projecteuler.net/problem=21
//
// ============================================================================

const max = 10_000

func sumOfDivisors(n int) int {
	var sum = 0
	for i := 1; i < n; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	return sum
}

func calculate() int {
	var sum = 0
	for a := 0; a < max; a++ {
		n := sumOfDivisors(a)
		b := sumOfDivisors(n)
		if a != n && a == b {
			sum += a
		}
	}
	return sum
}

func Solve() int {
	return calculate()
}
