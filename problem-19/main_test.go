package main

import (
	"testing"
)

var testCases = []struct {
	name   string
	dateS  string
	dateE  string
	output int
}{
	{"Basic", "1900-12-01 12:00:00", "1901-12-01 12:00:00", 2},
	{"Final", "1900-12-01 12:00:00", "2000-12-01 12:00:00", 171},
}

func TestSumOfMultiples(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			want := testCase.output
			have := getNumberOfSundaysBetwenDates(testCase.dateS, testCase.dateE)
			if want != have {
				t.Fatalf("Expected %d, Provided %d", want, have)
			}
		})
	}
}
