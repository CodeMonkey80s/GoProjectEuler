package main

import (
	"fmt"
	"time"

	"github.com/golang-module/carbon"
)

// Problem 19
//
// URL: https://projecteuler.net/problem=19
//
// ============================================================================

const dateS string = "1900-12-01 12:00:00"
const dateE string = "2000-12-01 12:00:00"

func getNumberOfSundaysBetwenDates(s1 string, s2 string) int {
	var numberOfSundays int = 0
	dateTimeS := carbon.Parse(s1)
	dateTimeE := carbon.Parse(s2)
	for {
		isSunday := false
		dateTimeS = dateTimeS.AddMonth()
		day := dateTimeS.Format("j")
		if dateTimeS.IsSunday() {
			isSunday = true
		}
		if isSunday && day == "1" {
			numberOfSundays++
		}
		if dateTimeS.Carbon2Time().Unix() >= dateTimeE.Carbon2Time().Unix() {
			break
		}
	}

	return numberOfSundays
}

func solve() int {
	return getNumberOfSundaysBetwenDates(dateS, dateE)
}

func main() {
	var t1 time.Time = time.Now()
	v := solve()
	fmt.Println(v)
	var t2 time.Duration = time.Since(t1)
	fmt.Printf("Time: %dms\n", t2.Milliseconds())
}
