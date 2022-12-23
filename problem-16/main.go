package main

import (
	"fmt"
	"math/big"
	"strconv"
	"time"
)

// Problem 16
//
// URL: https://projecteuler.net/problem=16
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
	var i, e = big.NewInt(2), big.NewInt(1000)
	i.Exp(i, e, nil)
	sumAsString := i.String()
	return sumOfDigits(sumAsString)
}

func main() {
	t1 := time.Now()
	v := solve()
	fmt.Println(v)
	t2 := time.Since(t1)
	fmt.Printf("Time: %dms\n", t2.Milliseconds())
}
