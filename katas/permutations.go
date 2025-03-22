package katas

import (
	"strings"
)

// Kata: So Many Permutations! <4kyu>
// https://www.codewars.com/kata/5254ca2719453dcc0b00027d
func Permutate(s string) []string {
	res := make([]string, 0)
	inner := strings.Split(s, "")
	if len(inner) == 1 {
		return inner
	}
	outer := strings.Split(s, "")

	sb := strings.Builder{}
	for o_idx, out := range outer {
		sb.WriteString(out)
		for i_idx, in := range inner {
			if o_idx == i_idx {
				continue
			}
			sb.WriteString(in)
			if sb.Len() == len(s) {
				res = append(res, sb.String())
				sb = strings.Builder{}
			}
		}
	}
	return res
}
