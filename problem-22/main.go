package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sort"
	"time"
)

// Problem 22
//
// URL: https://projecteuler.net/problem=22
//
// ============================================================================

const filename string = "names.txt"

var names = []string{}

func loadNames() {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	reader := csv.NewReader(f)
	names, _ = reader.Read()
	sort.Strings(names)
}

func valueOfName(n string) int {
	var value int = 0
	for i := 0; i < len(n); i++ {
		value += int(n[i]) - 64
	}
	return value
}

func solve() int {
	var sum int = 0
	loadNames()
	for i, v := range names {
		index := i + 1
		value := valueOfName(v)
		score := index * value
		sum += score
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
