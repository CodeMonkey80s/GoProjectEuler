package main

import (
	"fmt"
	"time"
)

// Problem 21
//
// URL: https://projecteuler.net/problem=21
//
// ============================================================================

const max = 10_000

func sumOfDivisors(n int) int {
	var sum int = 0
	for i := 1; i < n; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	return sum
}

func solve() int {
	var sum int = 0
	for a := 0; a < max; a++ {
		n := sumOfDivisors(a)
		b := sumOfDivisors(n)
		if a != n && a == b {
			sum += a
		}
	}
	return sum
}

func main() {
	t1 := time.Now()
	v := solve()
	fmt.Println(v)
	t2 := time.Since(t1)
	fmt.Printf("Time: %dms\n", t2.Milliseconds())
}
