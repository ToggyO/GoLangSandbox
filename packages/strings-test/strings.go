package strings_test

import "fmt"

func itoa(i int) (s string) {
	if i == 0 {
		return "0"
	}

	var isNegative bool
	if i < 0 {
		isNegative = true
		i *= -1
	}

	for i > 0 {
		pos := i % 10
		char := pos + '0'
		s = string(char) + s
		i /= 10
	}

	if isNegative {
		s = "-" + s
	}
	return s
}

func RunStrings() {
	type pair struct {
		i int
		s string
	}
	test := []pair{
		{0, "0"},
		{22, "22"},
		{32432523, "32432523"},
		{-3, "-3"},
	}
	for _, t := range test {
		if t.s == itoa(t.i) {
			fmt.Printf("%d - %s\n", t.i, "OK")
		} else {
			fmt.Printf("%d - %s\n", t.i, "FAIL")
		}
	}
}
