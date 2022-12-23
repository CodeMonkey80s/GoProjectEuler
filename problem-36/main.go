package main

import (
	"fmt"
	"strconv"
	"time"
)

// Problem 36
//
// URL: https://projecteuler.net/problem=36
//
// ============================================================================

const max int = 1_000_000

// NOTE: There is no strings.Reverse() function in the standard Go library
// URL: https://stackoverflow.com/questions/1752414/how-to-reverse-a-string-in-go
func Reverse(in string) (out string) {
	for _, r := range in {
		out = string(r) + out
	}
	return
}

func isPalindroneDecimal(n int) bool {
	str := strconv.Itoa(n)
	strReversed := Reverse(str)
	if str == strReversed {
		return true
	}
	return false
}

func isPalindroneBinary(n int) bool {
	str := fmt.Sprintf("%b", n)
	strReversed := Reverse(str)
	if str[0] != '0' && strReversed[0] != 0 && str == strReversed {
		return true
	}
	return false
}

func solve() int {
	var sum int = 0
	for i := 0; i <= max; i++ {
		pDec := isPalindroneDecimal(i)
		pBin := isPalindroneBinary(i)
		if pDec == true && pBin == true {
			sum += i
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
