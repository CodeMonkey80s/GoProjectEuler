package problem010

import (
	"testing"
)

var testCases = []struct {
	name   string
	input  int
	output int
}{
	{"Basic", 17, 41},
	{"Basic", 1000, 76127},
}

func TestSumOfEvenValuedTerms(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			want := testCase.output
			have := findSumOfPrimesBelowNumber(testCase.input)
			if want != have {
				t.Fatalf("Expected %d, Provided %d", want, have)
			}
		})
	}
}
