package problem017

// Problem 17
//
// URL: https://projecteuler.net/problem=17
//
// If the numbers 1 to 5 are written out in words: one, two, three, four, five,
// then there are 3 + 3 + 5 + 4 + 4 = 19 letters used in total.
//
// If all the numbers from 1 to 1000 (one thousand) inclusive were written out
// in words, how many letters would be used?
//
// NOTE: Do not count spaces or hyphens. For example, 342 (three hundred and
// forty-two) contains 23 letters and 115 (one hundred and fifteen) contains
// 20 letters. The use of "and" when writing out numbers is in compliance with
// British usage.
//
// ============================================================================

var NumberToWord = map[int]string{
	1:    "one",
	2:    "two",
	3:    "three",
	4:    "four",
	5:    "five",
	6:    "six",
	7:    "seven",
	8:    "eight",
	9:    "nine",
	10:   "ten",
	11:   "eleven",
	12:   "twelve",
	13:   "thirteen",
	14:   "fourteen",
	15:   "fifteen",
	16:   "sixteen",
	17:   "seventeen",
	18:   "eighteen",
	19:   "nineteen",
	20:   "twenty",
	30:   "thirty",
	40:   "forty",
	50:   "fifty",
	60:   "sixty",
	70:   "seventy",
	80:   "eighty",
	90:   "ninety",
	100:  "hundred",
	1000: "thousand",
}

const max = 1000

func convertFrom1To99(n int) (str string) {
	if n < 20 {
		return NumberToWord[n]
	}
	r := n % 10
	if r == 0 {
		return NumberToWord[n]
	} else {
		return NumberToWord[n-r] + NumberToWord[r]
	}
}

func convertFrom100To999(n int) (str string) {
	q := n / 100
	r := n % 100
	str = NumberToWord[q] + NumberToWord[100]
	if r == 0 {
		return
	} else {
		str = str + "and" + convertFrom1To99(r)
	}
	return
}

func convert(n int) (str string) {
	if n < 100 {
		str = convertFrom1To99(n)
	} else if n >= 100 && n <= 999 {
		str = convertFrom100To999(n)
	} else {
		return "one" + NumberToWord[n]
	}
	return str
}

func calculate() int {
	sum := 0
	letters := 0
	for i := 1; i <= max; i++ {
		text := convert(i)
		letters = len(text)
		sum += letters
	}
	return sum
}

func Solve() int {
	return calculate()
}
