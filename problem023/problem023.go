package problem023

// Problem 23
//
// URL: https://projecteuler.net/problem=23
//
// ============================================================================

const max int = 28123

var abundantNumbers = make([]int, 0, max)
var nonAbSums = make(map[int]int)

func isAbundantNumber(n int) bool {
	var sum = 1
	for x := 2; x < (n/2)+1; x++ {
		if n%x == 0 {
			sum += x
		}
	}
	return sum > n
}

func calculate() int {
	var sum = 0
	for i := 12; i < max; i++ {
		state := isAbundantNumber(i)
		if state == true {
			abundantNumbers = append(abundantNumbers, i)
		}
	}
	for _, yv := range abundantNumbers {
		for _, xv := range abundantNumbers {
			abSum := yv + xv
			if abSum > max {
				break
			} else {
				nonAbSums[abSum] = abSum
			}
		}
	}
	for n := 0; n < max; n++ {
		_, ok := nonAbSums[n]
		if ok == true {
			continue
		}
		sum += n
	}
	return sum
}

func Solve() int {
	return calculate()
}
