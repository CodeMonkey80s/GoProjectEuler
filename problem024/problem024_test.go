package problem024

import "testing"

var testCases = []struct {
	name   string
	digits uint64
	max    uint64
	output string
}{
	{"Basic", 3, 6, "210"},
	{"Basic", 10, 980_000, "2735084196"},
	{"Final", 10, 1_000_000, "2783915460"},
}

func TestSumOfMultiples(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			want := testCase.output
			have := calculate(testCase.max, testCase.digits)
			if want != have {
				t.Fatalf("Expected %s, Provided %s", want, have)
			}
		})
	}
}
