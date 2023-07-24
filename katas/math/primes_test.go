package math

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_DecompGoRout(t *testing.T) {

	res := DecompGoRout(3) // !3 = 6
	assert.Equal(t, "2 * 3", res)

	res = DecompGoRout(11)
	assert.Equal(t, "2^8 * 3^4 * 5^2 * 7 * 11", res)
}

func Test_Decomp(t *testing.T) {

	res := Decomp(3) // !3 = 6
	assert.Equal(t, "2 * 3", res)

	res = Decomp(11)
	assert.Equal(t, "2^8 * 3^4 * 5^2 * 7 * 11", res)
}

func Test_DecompV2(t *testing.T) {

	res := DecompV2(3) // !3 = 6
	assert.Equal(t, "2 * 3", res)

	res = DecompV2(12)
	assert.Equal(t, "2^10 * 3^5 * 5^2 * 7 * 11", res)

	// res := DecompV2(22)
	// assert.Equal(t, "2^19 * 3^9 * 5^4 * 7^3 * 11^2 * 13 * 17 * 19", res)

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

	res = getPrimes(120)
	assert.Equal(t, 30, len(res))
	assert.ElementsMatch(t, []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113}, res)
}

func Test_getPrimesV2(t *testing.T) {
	res := getPrimesV2(2)
	assert.ElementsMatch(t, []int{2}, res)

	res = getPrimesV2(5)
	assert.ElementsMatch(t, []int{2, 3, 5}, res)

	res = getPrimesV2(6)
	assert.ElementsMatch(t, []int{2, 3, 5}, res)

	res = getPrimesV2(7)
	assert.ElementsMatch(t, []int{2, 3, 5, 7}, res)

	res = getPrimesV2(120)
	assert.Equal(t, 30, len(res))
	assert.ElementsMatch(t, []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113}, res)
}

func Test_getPrimes_Laufzeit(t *testing.T) {
	start := time.Now()
	getPrimes(100000000)
	fmt.Printf("getPrimes() dauerte %v ms\n", time.Since(start).Milliseconds())
	start = time.Now()
	getPrimesV2(100000000)
	fmt.Printf("getPrimesV2() dauerte %v ms\n", time.Since(start).Milliseconds())
}
func Benchmark_getPrimes(b *testing.B) {
	testnums := []int{1000, 1000000, 10000000}

	for _, p := range testnums {

		b.Run(fmt.Sprintf("Prime %v", p), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				getPrimes(p)
			}
		})

		b.Run(fmt.Sprintf("PrimeV2 %v", p), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				getPrimesV2(p)
			}
		})
	}
}

func Benchmark_decomp(b *testing.B) {
	testnums := []int{5, 6, 7, 8, 9, 10}

	for _, p := range testnums {

		b.Run(fmt.Sprintf("Decomp !%v", p), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Decomp(p)
			}
		})

		b.Run(fmt.Sprintf("DecompGoRout !%v", p), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				DecompGoRout(p)
			}
		})

		b.Run(fmt.Sprintf("DecompV2 !%v", p), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				DecompV2(p)
			}
		})
	}
}

func Test_filterPrimes(t *testing.T) {
	res := filterPrimes([]int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, []int{2, 3, 5, 7, 11})
	assert.ElementsMatch(t, []int{2, 3, 5, 7, 11}, res)
}
