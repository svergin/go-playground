package math

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	count int
	v     version
}

func Test_getDividablesCount_simple(t *testing.T) {
	res := GetDividablesCount(V1, 6, 11, 2)
	assert.Equal(t, 3, res)

	res = GetDividablesCount(V2, 6, 11, 2)
	assert.Equal(t, 3, res)

	res = GetDividablesCount(V3, 6, 11, 2)
	assert.Equal(t, 3, res)
}

func Test_getDividablesCount_mid_complex(t *testing.T) {
	res := GetDividablesCount(V1, 1, 1000, 2)
	assert.Equal(t, 500, res)
	res = GetDividablesCount(V2, 1, 1000, 2)
	assert.Equal(t, 500, res)
	res = GetDividablesCount(V3, 1, 1000, 2)
	assert.Equal(t, 500, res)
}

func Test_getDividablesCount_Laufzeit(t *testing.T) {

	tc := []testcase{
		{count: 100000000, v: V1},
		{count: 100000000, v: V2},
		{count: 100000000, v: V3},
		{count: 1000000000, v: V1},
		{count: 1000000000, v: V2},
		{count: 1000000000, v: V3},
	}

	for _, c := range tc {
		start := time.Now()
		GetDividablesCount(c.v, 1, c.count, 3)
		fmt.Printf("getDividablesCount(%v, 1, %v, 3) dauerte %v ms\n", c.v, c.count, time.Since(start).Milliseconds())
	}

}

func Benchmark_getDividablesCount(b *testing.B) {
	tc := []testcase{
		{count: 10000000, v: V1},
		{count: 10000000, v: V2},
		{count: 10000000, v: V3},
		{count: 100000000, v: V1},
		{count: 100000000, v: V2},
		{count: 100000000, v: V3},
		{count: 1000000000, v: V1},
		{count: 1000000000, v: V2},
		{count: 1000000000, v: V3},
	}

	for _, c := range tc {
		b.Run(fmt.Sprintf("Version: %v, Zahl %v", c.v, c.count), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				GetDividablesCount(c.v, 1, c.count, 3)
			}
		})
	}
}
