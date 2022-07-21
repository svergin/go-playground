package katas

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Maps_Should_double_up_int_vals_in_array(t *testing.T) {
	input := []int{1, 2, 3, 4}
	actual := Maps(input)

	assert.Equal(t, actual[0], 2)
	assert.Equal(t, actual[1], 4)
	assert.Equal(t, actual[2], 6)
	assert.Equal(t, actual[3], 8)

}

func Test_Square_and_Sum_should_return_correct_result(t *testing.T) {
	input := []int{1, 2}
	actual := SquareSum(input)

	assert.Equal(t, 5, actual)

	input = []int{0, 3, 4, 5}
	actual = SquareSum(input)

	assert.Equal(t, 50, actual)

	input = []int{}
	actual = SquareSum(input)

	assert.Equal(t, 0, actual)

}

func Test_Spaces_removed_from_string_successfully(t *testing.T) {
	actual := NoSpace("Das ist ein Test.")

	assert.Equal(t, "DasisteinTest.", actual)

}

func Test_Positive_Numbers_were_correctly_summed_up(t *testing.T) {
	var toSum []int = []int{1, -2, 3, 4, 5}
	actual := PositiveSum(toSum)

	assert.Equal(t, 13, actual)
}

func Test_Herons_formula_calculated_the_correct_result(t *testing.T) {
	actual := Heron(3.00, 4.00, 5.00)
	assert.Equal(t, 6.00, actual)
}

func Test_GetMiddle_should_return_the_correct_string(t *testing.T) {
	actual := GetMiddle("Apfel")
	assert.Equal(t, "f", actual)
}
