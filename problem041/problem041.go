package problem041

// Problem 41
//
// URL: https://projecteuler.net/problem=41
//
// NOT SOLVED!
//
// ============================================================================

// https://en.wikipedia.org/wiki/Heap's_algorithm
// https://yourbasic.org/golang/generate-permutation-slice-string/
// https://stackoverflow.com/questions/30226438/generate-all-permutations-in-go

// func isPrimeNumber(n int) bool {
// 	for i := 2; i <= int(math.Floor(float64(n)/2)); i++ {
// 		if n%i == 0 {
// 			return false
// 		}
// 	}
// 	return bool(n > 1)
// }

// func arrayToInteger(array []int) int {
// 	var b bytes.Buffer
// 	for i := range array {
// 		b.WriteString(fmt.Sprintf("%d", array[i]))
// 	}
// 	n, _ := strconv.Atoi(b.String())
// 	return n
// }

// // Taken from: https://stackoverflow.com/a/55885234/1449403
// func generateIntPermutations(array []int, n int, result *[][]int) {
// 	if n == 1 {
// 		dst := make([]int, len(array))
// 		copy(dst, array[:])
// 		*result = append(*result, dst)
// 		n := arrayToInteger(dst)
// 		isPrime := isPrimeNumber(n)
// 		fmt.Printf("%#v = %v\n", n, isPrime)
// 		if isPrime {
// 			return
// 		}
// 	} else {
// 		for i := 0; i < n; i++ {
// 			generateIntPermutations(array, n-1, result)
// 			if n%2 == 0 {
// 				array[0], array[n-1] = array[n-1], array[0]
// 			} else {
// 				array[i], array[n-1] = array[n-1], array[i]
// 			}
// 		}
// 	}
// }

// func solve() int {
// 	var number int = 0
// 	var numbers = []int{4, 3, 2, 1, 0}
// 	// var numbers = []int{9, 8, 7, 6, 5, 3, 4, 3, 2, 1, 0}
// 	var result [][]int
// 	generateIntPermutations(numbers, len(numbers), &result)
// 	return number
// }

// func main() {
// 	t1 := time.Now()
// 	v := solve()
// 	fmt.Println(v)
// 	t2 := time.Since(t1)
// 	fmt.Printf("Time: %dms\n", t2.Milliseconds())
// }

func Solve() int {
	return 0
}
