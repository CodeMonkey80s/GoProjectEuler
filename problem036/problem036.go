package problem036

import (
	"fmt"
	"strconv"
)

// Problem 36
//
// URL: https://projecteuler.net/problem=36
//
// ============================================================================

const max int = 1_000_000

// Reverse
// NOTE: There is no strings.Reverse() function in the standard Go library
// URL: https://stackoverflow.com/questions/1752414/how-to-reverse-a-string-in-go
func Reverse(in string) (out string) {
	for _, r := range in {
		out = string(r) + out
	}
	return
}

func isPalindromeDecimal(n int) bool {
	str := strconv.Itoa(n)
	strReversed := Reverse(str)
	if str == strReversed {
		return true
	}
	return false
}

func isPalindromeBinary(n int) bool {
	str := fmt.Sprintf("%b", n)
	strReversed := Reverse(str)
	if str[0] != '0' && strReversed[0] != 0 && str == strReversed {
		return true
	}
	return false
}

func calculate() int {
	var sum = 0
	for i := 0; i <= max; i++ {
		pDec := isPalindromeDecimal(i)
		pBin := isPalindromeBinary(i)
		if pDec == true && pBin == true {
			sum += i
		}
	}
	return sum
}

func Solve() int {
	return calculate()
}
