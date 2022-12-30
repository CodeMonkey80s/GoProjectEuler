package problem002

import (
	"testing"
)

var testCases = []struct {
	name   string
	input  int
	output int
}{
	{"Basic", 10, 10},
	{"Basic", 100, 188},
	{"Final", 4_000_000, 4613732},
}

func TestSumOfEvenValuedTerms(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			want := testCase.output
			have := sumOfEvenValuedTerms(testCase.input)
			if want != have {
				t.Fatalf("Expected %d, Provided %d", want, have)
			}
		})
	}
}
