package problem004

import (
	"log"
	"strconv"
	"strings"
)

// Problem 4
//
// URL: https://projecteuler.net/problem=4
//
// A palindromic number reads the same both ways. The largest palindrome made
// from the product of two 2-digit numbers is 9009 = 91 × 99.
//
// Find the largest palindrome made from the product of two 3-digit numbers.
//
// ============================================================================

const digits int = 3

// Reverse
// NOTE: There is no strings.Reverse() function in the standard Go library
// URL: https://stackoverflow.com/questions/1752414/how-to-reverse-a-string-in-go
func Reverse(in string) (out string) {
	for _, r := range in {
		out = string(r) + out
	}
	return
}

func bounds(digits int) (int, int) {
	digits--
	if digits < 0 {
		log.Fatal("Number of digits can't be lower than zero!")
	}
	min, _ := strconv.Atoi("1" + strings.Repeat("0", digits))
	max, _ := strconv.Atoi("9" + strings.Repeat("9", digits))
	return min, max
}

func palindrome(digits int) uint64 {
	var largest uint64 = 0
	min, max := bounds(digits)
	for y := max; y > min; y-- {
		for x := max; x > min; x-- {
			n := uint64(x * y)
			str1 := strconv.FormatUint(n, 10)
			str2 := Reverse(str1)
			if str1 == str2 {
				if n > largest {
					largest = n
				}
			}
		}
	}
	return largest
}

func Solve() uint64 {
	return palindrome(digits)
}
