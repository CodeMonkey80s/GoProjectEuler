package main

import (
	"fmt"
	"time"
)

// Problem 28
//
// URL: https://projecteuler.net/problem=28
//
// ============================================================================

const size = 1001

var matrix = [size][size]int{}

func printMatrix() {
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			fmt.Printf("%d ", matrix[y][x])
		}
		fmt.Println()
	}
}

func fillMatrix() {
	var dy = []int{1, 0, -1, 0}
	var dx = []int{0, -1, 0, 1}
	var count = (size - 1) / 2
	var i int = 1
	var x int = count
	var y int = count
	var repeat = 0
	matrix[x][y] = 1
	for ring := 1; ring <= count; ring++ {
		repeat += 2
		y--
		x++
		for dir := 0; dir < 4; dir++ {
			for r := 0; r < repeat; r++ {
				y += dy[dir]
				x += dx[dir]
				i++
				matrix[y][x] = i
			}
		}
	}
}

func solve() int {
	fillMatrix()
	var sum int = 1
	var m = size / 2
	for i := 1; i <= m; i++ {
		sum += matrix[m-i][m-i]
		sum += matrix[m-i][m+i]
		sum += matrix[m+i][m-i]
		sum += matrix[m+i][m+i]
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
