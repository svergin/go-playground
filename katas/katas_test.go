package katas

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Maps_Should_double_up_int_vals_in_array(t *testing.T) {
	input := []int{1, 2, 3, 4}
	output := Maps(input)

	assert.Equal(t, output[0], 2)
	assert.Equal(t, output[1], 4)
	assert.Equal(t, output[2], 6)
	assert.Equal(t, output[3], 8)

}

func Test_Square_and_Sum_should_return_correct_result(t *testing.T) {
	input := []int{1, 2}
	output := SquareSum(input)

	assert.Equal(t, 5, output)

	input = []int{0, 3, 4, 5}
	output = SquareSum(input)

	assert.Equal(t, 50, output)

	input = []int{}
	output = SquareSum(input)

	assert.Equal(t, 0, output)

}
