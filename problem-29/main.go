package main

import (
	"fmt"
	"math/big"
	"time"
)

// Problem 29
//
// URL: https://projecteuler.net/problem=29
//
// ============================================================================

func solve() int {
	var a_min = 2
	var a_max = 100
	var b_min = 2
	var b_max = 100
	var numbers = make(map[string]string)
	var index int = 0
	for b := b_min; b <= b_max; b++ {
		for a := a_min; a <= a_max; a++ {
			var i, e = big.NewInt(int64(a)), big.NewInt(int64(b))
			i.Exp(i, e, nil)
			str := i.String()
			numbers[str] = i.String()
			index++
		}
	}
	return len(numbers)
}

func main() {
	t1 := time.Now()
	v := solve()
	fmt.Println(v)
	t2 := time.Since(t1)
	fmt.Printf("Time: %dms\n", t2.Milliseconds())
}
