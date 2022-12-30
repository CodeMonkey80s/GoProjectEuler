package problem020

import (
	"math/big"
	"strconv"
)

// Problem 20
//
// URL: https://projecteuler.net/problem=20
//
// ============================================================================

func sumOfDigits(s string) int {
	var sum = 0
	for i := 0; i < len(s); i++ {
		v, _ := strconv.Atoi(s[i : i+1])
		sum += v
	}
	return sum
}

func calculate() int {
	var i = big.Int{}
	i.MulRange(1, 100)
	var sumAsString = i.String()
	return sumOfDigits(sumAsString)
}

func Solve() int {
	return calculate()
}
