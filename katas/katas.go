package katas

import (
	"math"
	"strings"
)

//Kata: Beginner - Lost Without a Map
func Maps(x []int) []int {
	output := make([]int, len(x))

	for pos, v := range x {
		output[pos] = v * 2
	}
	return output
}

//Kata: Square(n) Sum
func SquareSum(numbers []int) int {
	var result int = 0

	for _, v := range numbers {
		result += v * v
	}
	return result
}

//Kata: Remove String Spaces
func NoSpace(word string) string {
	return strings.ReplaceAll(word, " ", "")
}

//Kata: Sum of positive
func PositiveSum(numbers []int) int {
	var result int = 0
	for _, v := range numbers {
		if v >= 0 {
			result += v
		}
	}
	return result
}

//Kata: Heron's formula
func Heron(a, b, c float64) (area float64) {
	s := (a + b + c) / 2
	area = math.Sqrt(s * (s - a) * (s - b) * (s - c))
	return area
}

//Kata: Get the Middle Character
func GetMiddle(s string) string {
	laenge := len(s)

	if laenge%2 == 1 {
		//TODO
	}

	return ""
}
