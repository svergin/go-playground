package math

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Decomp(t *testing.T) {

	// res := DecompGoRout(3) // !3 = 6
	// assert.Equal(t, "2^1 * 3^1", res)

	res := DecompGoRout(11)
	assert.Equal(t, "2^10 * 3^5 * 5^2 * 7 * 11", res)

}

func Test_getPrimes(t *testing.T) {
	// res := getPrimes(2)
	// assert.ElementsMatch(t, []int{2}, res)

	// res = getPrimes(5)
	// assert.ElementsMatch(t, []int{2, 3, 5}, res)

	// res = getPrimes(6)
	// assert.ElementsMatch(t, []int{2, 3, 5}, res)

	// res = getPrimes(7)
	// assert.ElementsMatch(t, []int{2, 3, 5, 7}, res)

	res := getPrimes(120)
	assert.Equal(t, 30, len(res))
	assert.ElementsMatch(t, []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113}, res)
}

func Benchmark_getPrimes(b *testing.B) {
	testnums := []int{1000, 1000000, 10000000}

	for _, p := range testnums {

		b.Run(fmt.Sprintf("Prime %v", p), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				getPrimes(p)
			}
		})
	}
}

func Test_filterPrimes(t *testing.T) {
	res := filterPrimes([]int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, []int{2, 3, 5, 7, 11})
	assert.ElementsMatch(t, []int{2, 3, 5, 7, 11}, res)
}
