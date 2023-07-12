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

func Test_GetMiddle_should_return_the_correct_string_uneven_character_count(t *testing.T) {
	actual := GetMiddle("Apfel")
	assert.Equal(t, "f", actual)
	actual = GetMiddle("Apfelkuchen")
	assert.Equal(t, "k", actual)
}

func Test_GetMiddle_should_return_the_correct_string_even_character_count(t *testing.T) {
	actual := GetMiddle("Obst")
	assert.Equal(t, "bs", actual)

}

func Test_GetCount_should_rreturn_the_correct_number_of_vowels(t *testing.T) {
	actual := GetCount("otto")
	assert.Equal(t, 2, actual)
	actual = GetCount("usambaraveilchen")
	assert.Equal(t, 7, actual)
}

func Test_HighAndLow_should_return_the_correct_result(t *testing.T) {
	actual := HighAndLow("1 2 3 4 5")
	assert.Equal(t, "5 1", actual)
	actual = HighAndLow("1 2 -3 4 5")
	assert.Equal(t, "5 -3", actual)
	actual = HighAndLow("1 9 3 4 -5")
	assert.Equal(t, "9 -5", actual)
}

func Test_Accum_should_produce_the_correct_output(t *testing.T) {
	actual := Accum("abcd")
	assert.Equal(t, "A-Bb-Ccc-Dddd", actual)

	actual = Accum("RqaEzty")
	assert.Equal(t, "R-Qq-Aaa-Eeee-Zzzzz-Tttttt-Yyyyyyy", actual)

	actual = Accum("cwAt")
	assert.Equal(t, "C-Ww-Aaa-Tttt", actual)
}

func Test_FindShortest_should_return_shortest_word(t *testing.T) {
	actual := FindShortest("Ein kurzer Satz.")
	assert.Equal(t, 3, actual)
	actual = FindShortest("Ein dickes Ei!")
	assert.Equal(t, 2, actual)
}

func Test_TwoToOne_should_return_correct_sorted_string_with_single_occurance_of_every_letter(t *testing.T) {
	actual := TwoToOne("aaxxbb", "yycccz")
	assert.Equal(t, "abcxyz", actual)

	actual = TwoToOne("aretheyhere", "yestheyarehere")
	assert.Equal(t, "aehrsty", actual)

}

func Test_Decomp(t *testing.T) {

	// res := DecompGoRout(3) // !3 = 6
	// assert.Equal(t, "2^1 * 3^1", res)

	res := DecompGoRout(11)
	assert.Equal(t, "2^10 * 3^5 * 5^2 * 7 * 11", res)

}

func Test_getPrimes(t *testing.T) {
	res := getPrimes(2)
	assert.ElementsMatch(t, []int{2}, res)

	res = getPrimes(5)
	assert.ElementsMatch(t, []int{2, 3, 5}, res)

	res = getPrimes(6)
	assert.ElementsMatch(t, []int{2, 3, 5}, res)

	res = getPrimes(7)
	assert.ElementsMatch(t, []int{2, 3, 5, 7}, res)
}

func Test_filterPrimes(t *testing.T) {
	res := filterPrimes([]int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, []int{2, 3, 5, 7, 11})
	assert.ElementsMatch(t, []int{2, 3, 5, 7, 11}, res)
}
