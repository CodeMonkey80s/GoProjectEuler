package main

import (
	"fmt"
	"math/big"
	"time"
)

// Problem 25
//
// URL: https://projecteuler.net/problem=25
//
// ============================================================================

// const maxLength int = 12
const maxLength int = 1000

func fibonacci(n *big.Int) *big.Int {
	var a = big.NewInt(0)
	var b = big.NewInt(1)
	var s = big.NewInt(0)
	var max = n
	var inc = big.NewInt(1)
	for i := new(big.Int).Set(s); i.Cmp(max) < 0; i.Add(i, inc) {
		c := a
		a = b
		b = c.Add(c, a)
	}
	return a
}

func solve(max int) *big.Int {
	var index = big.NewInt(0)
	var inc = big.NewInt(1)
	var i = big.NewInt(1)
	for {
		n := fibonacci(i)
		l := len(n.String())
		if l >= max {
			index = i
			break
		}
		i.Add(i, inc)
	}

	return index
}

func main() {
	t1 := time.Now()
	v := solve(maxLength)
	fmt.Println(v)
	t2 := time.Since(t1)
	fmt.Printf("Time: %dms\n", t2.Milliseconds())
}
