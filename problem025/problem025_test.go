package problem025

import (
	"math/big"
	"testing"
)

var testCasesFibonacci = []struct {
	name   string
	index  int64
	output int64
}{
	{"Basic", 0, 0},
	{"Basic", 1, 1},
	{"Basic", 2, 1},
	{"Basic", 3, 2},
	{"Basic", 4, 3},
	{"Basic", 5, 5},
	{"Basic", 6, 8},
	{"Basic", 7, 13},
	{"Basic", 12, 144},
}

func TestFibonacci(t *testing.T) {
	for _, testCase := range testCasesFibonacci {
		t.Run(testCase.name, func(t *testing.T) {
			want := big.NewInt(testCase.output)
			index := big.NewInt(testCase.index)
			have := fibonacci(index)
			if have.Cmp(want) != 0 {
				t.Fatalf("Expected %v, Provided %v", want, have)
			}
		})
	}
}

var testCasesSolve = []struct {
	name   string
	length int64
	output int64
}{
	{"Basic", 3, 12},
	{"Final", 1000, 4782},
}

func TestSolve(t *testing.T) {
	for _, testCase := range testCasesSolve {
		t.Run(testCase.name, func(t *testing.T) {
			want := big.NewInt(testCase.output)
			have := calculate(int(testCase.length))
			if have.Cmp(want) != 0 {
				t.Fatalf("Expected %v, Provided %v", want, have)
			}
		})
	}
}
