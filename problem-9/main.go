package main

import (
	"fmt"
	"math"
)

// Problem 9
//
// URL: https://projecteuler.net/problem=9
//
// ============================================================================

func findProduct() int {
	for y := 1; y <= 1000; y++ {
		for x := 1; x <= 1000; x++ {
			a := float64(x)
			b := float64(y)
			cs := math.Pow(a, 2) + math.Pow(b, 2)
			c := math.Sqrt(cs)
			if a+b+c == 1000 {
				return int(a * b * c)
			}
		}
	}
	return 0
}

func main() {
	v := findProduct()
	fmt.Println(v)
}
