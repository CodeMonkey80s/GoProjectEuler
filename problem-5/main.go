package main

import (
	"fmt"
	"time"
)

// Problem 5
//
// URL: https://projecteuler.net/problem=5
//
// 2520 is the smallest number that can be divided by each of the numbers
// from 1 to 10 without any remainder.
//
// What is the smallest positive number that is evenly divisible by
// all of the numbers from 1 to 20?
//
// ============================================================================

const min, max int = 1, 20

func evenlyDivisible(min, max int) int {
	var divisible = true
	var number = 1
	var final = 0
	for {
		for i := min; i <= max; i++ {
			if number%i != 0 {
				divisible = false
			}
		}
		if divisible {
			final = number
			break
		}
		number++
		// This is just a failsafe
		if number > 1_000_000_000 {
			break
		}
		divisible = true
	}
	return final
}

func main() {
	t1 := time.Now()
	v := evenlyDivisible(min, max)
	fmt.Println(v)
	t2 := time.Since(t1)
	fmt.Printf("Time: %dms\n", t2.Milliseconds())
}
