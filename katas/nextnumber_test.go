package katas

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNextBiggerNumber_12(t *testing.T) {
	result := NextBiggerNumber(12)
	assert.Equal(t, 21, result)
}
func TestNextBiggerNumber_513(t *testing.T) {
	result := NextBiggerNumber(513)
	assert.Equal(t, 531, result)
}
func TestNextBiggerNumber_2017(t *testing.T) {
	result := NextBiggerNumber(2017)
	assert.Equal(t, 2071, result)
}
func TestNextBiggerNumber_764538888(t *testing.T) {
	result := NextBiggerNumber(764538888)
	assert.Equal(t, 764583888, result)
}
func TestNextBiggerNumber_111(t *testing.T) {
	result := NextBiggerNumber(111)
	assert.Equal(t, -1, result)
}
