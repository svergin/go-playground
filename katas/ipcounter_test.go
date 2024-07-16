package katas

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvert(t *testing.T) {
	ip, err := convert("10.0.1.233")
	if assert.NoError(t, err) {
		assert.Equal(t, int64(10), ip.first)
		assert.Equal(t, int64(0), ip.second)
		assert.Equal(t, int64(1), ip.third)
		assert.Equal(t, int64(233), ip.fourth)
	}
}

func TestIPCount_Erfolg(t *testing.T) {
	type testcase struct {
		startip       string
		endip         string
		expectedCount int64
	}

	cases := []testcase{
		{"10.0.0.0", "10.0.0.50", int64(50)},
		{"10.0.0.0", "10.0.1.0", int64(256)},
		{"20.0.0.10", "20.0.1.0", int64(246)},
		{"20.0.0.10", "20.0.1.30", int64(276)},
	}

	for _, c := range cases {
		res, err := CountIPs(c.startip, c.endip)
		assert.NoError(t, err)
		assert.Equal(t, c.expectedCount, res)
	}

}
