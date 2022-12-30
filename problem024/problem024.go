package problem024

import (
	"strconv"
)

// Problem 24
//
// URL: https://projecteuler.net/problem=24
//
// HELP: https://stemhash.com/efficient-permutations-in-lexicographic-order/
//
// ============================================================================

const max = 1_000_000
const digits = 10

func factorial(n uint64) uint64 {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

func calculate(max uint64, digits uint64) string {
	var numbers = make(map[uint64]uint64)
	var i uint64 = 0
	for i = 0; i < 10; i++ {
		numbers[i] = i
	}
	var next = max - 1
	var n = digits - 1
	var s = ""
	for {
		f := factorial(n)
		q := next / f
		r := next % f
		d := numbers[q]
		for i = 0; i < 10-1; i++ {
			if numbers[i] >= numbers[q] {
				numbers[i] = numbers[i+1]
			}
		}
		s += strconv.Itoa(int(d))
		next = r
		n--
		if n == 0 {
			break
		}
	}
	s += strconv.Itoa(int(numbers[0]))
	return s
}

func Solve() string {
	return calculate(max, digits)
}
