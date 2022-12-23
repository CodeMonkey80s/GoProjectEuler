package main

import (
	"fmt"
	"math"
	"time"
)

// Problem 7
//
// URL: https://projecteuler.net/problem=7
//
// By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13,
// we can see that the 6th prime is 13.
//
// What is the 10 001st prime number?
//
// ============================================================================

const position int = 10_001

func isPrimeNumber(n int) bool {
	for i := 2; i <= int(math.Floor(float64(n)/2)); i++ {
		if n%i == 0 {
			return false
		}
	}
	return bool(n > 1)
}

func findPrimeNumberAtPosition(position int) int {
	var counter int = 0
	var i = 2
	var final = 0
	for {
		v := isPrimeNumber(i)
		if v == true {
			final = i
			counter++
		}
		if counter == position {
			break
		}
		i++
	}
	return final
}

func main() {
	t1 := time.Now()
	v := findPrimeNumberAtPosition(position)
	fmt.Println(v)
	t2 := time.Since(t1)
	fmt.Printf("Time: %dms\n", t2.Milliseconds())
}
