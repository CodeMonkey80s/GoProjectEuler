package problem022

import (
	"encoding/csv"
	"log"
	"os"
	"sort"
)

// Problem 22
//
// URL: https://projecteuler.net/problem=22
//
// ============================================================================

const filename string = "names.txt"

var names []string

func loadNames() {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatalf("Can't open file %s!\n", filename)
		}
	}(f)
	reader := csv.NewReader(f)
	names, _ = reader.Read()
	sort.Strings(names)
}

func valueOfName(n string) int {
	var value = 0
	for i := 0; i < len(n); i++ {
		value += int(n[i]) - 64
	}
	return value
}

func calculate() int {
	var sum = 0
	loadNames()
	for i, v := range names {
		index := i + 1
		value := valueOfName(v)
		score := index * value
		sum += score
	}
	return sum
}

func Solve() int {
	return calculate()
}
