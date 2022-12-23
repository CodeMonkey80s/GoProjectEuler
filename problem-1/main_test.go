package main

import (
	"testing"
)

var testCases = []struct {
	name   string
	input  int
	output int
}{
	{"Basic", 10, 23},
	{"Basic", 100, 2318},
	{"Final", 1000, 233168},
}

func TestSumOfMultiples(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			want := testCase.output
			have := sumOfMultiples(testCase.input)
			if want != have {
				t.Fatalf("Expected %d, Provided %d", want, have)
			}
		})
	}
}
