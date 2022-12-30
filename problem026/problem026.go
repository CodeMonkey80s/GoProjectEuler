package problem026

// Problem 26
//
// URL: https://projecteuler.net/problem=26
//
// NOT SOLVED!
//
// ============================================================================

// Process numbers 1 to 10
// Obtain fraction part rx "1/6" is "0.1(6)"
// Identify recurring pattern

// const max int = 1000

// func recurringDigit(s string) int {
// 	var longest int = 0
// 	s = strings.TrimPrefix(s, "0.")
// 	for i := 0; i <= 9; i++ {
// 		r := strings.Count(s, strconv.Itoa(i))
// 		// fmt.Println(i, r)
// 		if r > longest {
// 			longest = r
// 		}
// 	}
// 	return longest
// }

// func longestRepeatingSequence(s1 string, s2 string) string {
// 	var n int = int(math.Min(float64(len(s1)), float64(len(s2))))
// 	for i := 0; i < n; i++ {
// 		if s1[i] != s2[i] {
// 			return s1[0:i]
// 		}
// 	}
// 	return s1[0:n]
// }

// func find(s string) int {
// 	var n int = len(s)
// 	var lrs string = ""
// 	for i := 0; i < n; i++ {
// 		for j := i + 1; j < n; j++ {
// 			x := longestRepeatingSequence(s[i:n], s[j:n])
// 			if len(x) > len(lrs) {
// 				lrs = x
// 			}
// 		}
// 	}
// 	fmt.Println("LRS", lrs, len(lrs))
// 	return len(lrs)
// }

// func findLengthOfRecurringPattern(s string) int {
// 	var v int = 0
// 	var l int = 0
// 	var n int = len(s)
// 	for i := 0; i <= n; i++ {
// 		for j := n; j > i; j-- {
// 			sub := s[i:j]
// 			b := strings.Count(s, sub)
// 			// fmt.Println(i, j, sub, b)
// 			if b >= 2 && len(sub) > l && b > v {
// 				v = b
// 				l = len(sub)
// 				// fmt.Println("###", b, sub, l)
// 				repStr := strings.Repeat(sub[0:1], l)
// 				if repStr == sub {
// 					l = 1
// 					// fmt.Println("L", l)
// 					goto end
// 				}
// 			}
// 		}
// 	}
// end:
// 	return l
// }

// func primeSieve(n float64) {
// 	sieve := 1 * (n / 2)
// 	for i := in range(3, int(n**0.5)+1,2) {
// 		if sieve[i /2] {
// 			sieve[i*i/2:i] = false * ((n-i*i-1)/(2*i)+1)
// 		}
// 	}
// 	return [2] + [2*i+1 for i in range(1, n/2) if sieve[i]]
// }

func calculate() int {
	var best = 0
	// var length int = 0
	// var maxLength int = 0
	// var limit int = 0
	// for i := 7; i < max; i++ {
	// 	length = 0
	// 	value := 10 % i
	// 	limit = i - 1
	// 	fmt.Println(i, value, limit, best)
	// 	for value != 1 && length == maxLength && length <= limit {
	// 		maxLength = length
	// 		best = i
	// 	}
	// }
	return best
}

func Solve() int {
	return calculate()
}
