package main

import (
	"fmt"
	"time"
)

// Problem 14
//
// URL: https://projecteuler.net/problem=14
//
// ============================================================================

const max int = 1_000_000

func calculateNumberOfTerms(n int) int {
	var i int = 1
	for {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = (3 * n) + 1
		}
		if n == 1 {
			i++
			break
		}
		i++
	}

	return i
}

func solve() int {
	var number int = 0
	var longest int = 0
	var i int = 1
	for {
		num := calculateNumberOfTerms(i)
		if num > longest {
			longest = num
			number = i
		}
		i++
		if i > max-1 {
			break
		}
	}
	return number
}

func main() {
	t1 := time.Now()
	v := solve()
	fmt.Println(v)
	t2 := time.Since(t1)
	fmt.Printf("Time: %dms\n", t2.Milliseconds())
}
