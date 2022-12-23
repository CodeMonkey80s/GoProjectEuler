package main

import (
	"fmt"
	"math"
	"time"
)

const maxDivisors int = 500

func countFactors(n int) int {
	var count int = 0
	var sqrt int = int(math.Sqrt(float64(n)))
	for i := 1; i <= sqrt; i++ {
		if n%i == 0 {
			count += 2
		}
		if sqrt*sqrt == n {
			count--
		}
	}
	return count
}

func findHighlyDivisibleTriangularNumber() int {
	var n int = 0
	var i int = 0
	var factors int = 0
	for {
		factors = countFactors(n)
		if factors > maxDivisors {
			break
		}
		n += i
		i++
	}
	return n
}

func main() {
	t1 := time.Now()
	v := findHighlyDivisibleTriangularNumber()
	fmt.Println(v)
	t2 := time.Since(t1)
	fmt.Printf("Time: %dms\n", t2.Milliseconds())
}
