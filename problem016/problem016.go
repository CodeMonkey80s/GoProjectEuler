package problem016

import (
	"math/big"
	"strconv"
)

// Problem 16
//
// URL: https://projecteuler.net/problem=16
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
	var i, e = big.NewInt(2), big.NewInt(1000)
	i.Exp(i, e, nil)
	sumAsString := i.String()
	return sumOfDigits(sumAsString)
}

func Solve() int {
	return calculate()
}
