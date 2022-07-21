package katas

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
