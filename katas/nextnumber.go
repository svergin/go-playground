package katas

import (
	"strconv"
	"strings"
)

// Kata: Next bigger number with the same digits <4 kyu>
// https://www.codewars.com/kata/55983863da40caa2c900004e
func NextBiggerNumber(n int) int {

	numString := strconv.Itoa(n)
	digits := make([]int, len(numString))
	res := make([]int, len(numString))
	for idx, n := range numString {
		d, err := strconv.ParseInt(string(n), 10, 0)
		if err != nil {
			panic(err)
		}
		digits[idx] = int(d)
	}
	ready := false
	for i := len(digits) - 1; i >= 0; i-- {
		if i == 0 && !ready {
			return -1
		}
		if ready {
			res[i] = digits[i]
			continue
		}
		if digits[i] > digits[i-1] {
			res[i-1] = digits[i]
			if i >= 1 {
				res[i] = digits[i-1]
				ready = true
				i--
			}
		} else {
			res[i] = digits[i]
		}
	}

	b := strings.Builder{}
	for _, v := range res {
		b.WriteString(strconv.Itoa(v))
	}
	r, err := strconv.ParseUint(b.String(), 10, 32)
	if err != nil {
		panic(err)
	}
	if uint(n) == uint(r) {
		return -1
	}
	return int(r)
}
