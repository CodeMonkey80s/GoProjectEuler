package main

import (
	"fmt"
)

// Problem 2
//
// URL: https://projecteuler.net/problem=2
//
// Each new term in the Fibonacci sequence is generated by adding the previous
// two terms. By starting with 1 and 2, the first 10 terms will be:
// 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...
//
// By considering the terms in the Fibonacci sequence whose values do not
// exceed four million, find the sum of the even-valued terms.
//
// ============================================================================

// const max int = 4_000_000
const max int = 100

func fibonacci(n int) int {
	var a = 0
	var b = 1
	var max = n
	for i := 0; i < max; i++ {
		c := a
		a = b
		b = c + a
	}
	return a
}

func sumOfEvenValuedTerms(max int) int {
	sum := 0
	for i := 0; i <= max; i++ {
		v := fibonacci(i)
		if v%2 == 0 {
			sum += v
		}
		if v >= max {
			break
		}
	}
	return sum
}

func main() {
	fmt.Println(sumOfEvenValuedTerms(max))
}
