package katas

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPermutations_a(t *testing.T) {
	a := "a"
	result := Permutate(a)
	assert.Equal(t, 1, len(result))
	assert.Contains(t, result, a)
}
func TestPermutations_ab(t *testing.T) {
	ab := "ab"
	result := Permutate(ab)
	assert.Equal(t, 2, len(result))
	assert.Contains(t, result, ab)
	assert.Contains(t, result, "ba")
}
func TestPermutations_abc(t *testing.T) {
	abc := "abc"
	result := Permutate(abc)
	assert.Equal(t, 6, len(result))
	assert.Contains(t, result, abc)
	assert.Contains(t, result, "acb")
	assert.Contains(t, result, "bac")
	assert.Contains(t, result, "bca")
	assert.Contains(t, result, "cab")
	assert.Contains(t, result, "cba")
}
func TestPermutations_aabb(t *testing.T) {
	aabb := "aabb"
	result := Permutate(aabb)
	assert.Equal(t, 6, len(result))
	assert.Contains(t, result, aabb)
	assert.Contains(t, result, "abab")
	assert.Contains(t, result, "abba")
	assert.Contains(t, result, "baab")
	assert.Contains(t, result, "baba")
	assert.Contains(t, result, "bbaa")
}
