package main

import (
	"fmt"
	"math/big"
	"strconv"
	"time"
)

// Problem 20
//
// URL: https://projecteuler.net/problem=20
//
// ============================================================================

func sumOfDigits(s string) int {
	var sum int = 0
	for i := 0; i < len(s); i++ {
		v, _ := strconv.Atoi(s[i : i+1])
		sum += v
	}
	return sum
}

func solve() int {
	var i = big.Int{}
	i.MulRange(1, 100)
	var sumAsString string = i.String()
	return sumOfDigits(sumAsString)
}

func main() {
	t1 := time.Now()
	v := solve()
	fmt.Println(v)
	t2 := time.Since(t1)
	fmt.Printf("Time: %dms\n", t2.Milliseconds())
}
