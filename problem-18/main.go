package main

import (
	"fmt"
	"time"
)

// Problem 18
//
// URL: https://projecteuler.net/problem=18
//
// HELP: https://stackoverflow.com/a/29472318/1449403
//
// ============================================================================

/*
const msx = 4
const msy = 4

var matrix = [msy][msx]int{
	{3, 0, 0, 0},
	{7, 4, 0, 0},
	{2, 4, 6, 0},
	{8, 5, 9, 3},
}
*/

const msx = 15
const msy = 15

var matrix = [msy][msx]int{
	{75, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{95, 64, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{17, 47, 82, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{18, 35, 87, 10, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{20, 4, 82, 47, 65, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{19, 1, 23, 75, 3, 34, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{88, 2, 77, 73, 7, 63, 67, 0, 0, 0, 0, 0, 0, 0, 0},
	{99, 65, 4, 28, 6, 16, 70, 92, 0, 0, 0, 0, 0, 0, 0},
	{41, 41, 26, 56, 83, 40, 80, 70, 33, 0, 0, 0, 0, 0, 0},
	{41, 48, 72, 33, 47, 32, 37, 16, 94, 29, 0, 0, 0, 0, 0},
	{53, 71, 44, 65, 25, 43, 91, 52, 97, 51, 14, 0, 0, 0, 0},
	{70, 11, 33, 28, 77, 73, 17, 78, 39, 68, 17, 57, 0, 0, 0},
	{91, 71, 52, 38, 17, 14, 91, 43, 58, 50, 27, 29, 48, 0, 0},
	{63, 66, 4, 68, 89, 53, 67, 30, 73, 16, 69, 87, 40, 31, 0},
	{4, 62, 98, 27, 23, 9, 70, 98, 73, 93, 38, 53, 60, 4, 23},
}

func calculateSumFromTopToBottom(m [msy][msx]int) int {
	var sum int = 0

	for y := msy - 1; y >= 1; y-- {
		for x := 0; x < msx-1; x++ {
			v1 := matrix[y][x]
			v2 := matrix[y][x+1]
			v1 = v1 + matrix[y-1][x]
			v2 = v2 + matrix[y-1][x]
			if v1 > v2 {
				matrix[y-1][x] = v1
				sum = v1
				if x == 0 && y == 1 {
					return v1
				}
			}
			if v2 > v1 {
				matrix[y-1][x] = v2
				if x == 0 && y == 1 {
					return v2
				}
			}
		}
	}

	return sum
}

func solve() int {
	return calculateSumFromTopToBottom(matrix)
}

func main() {
	var t1 time.Time = time.Now()
	v := solve()
	fmt.Println(v)
	var t2 time.Duration = time.Since(t1)
	fmt.Printf("Time: %dms\n", t2.Milliseconds())
}
