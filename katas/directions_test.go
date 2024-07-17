package katas

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDirReduc_Simple_Erfolg(t *testing.T) {
	res := DirReduc([]direction{North, South, East, West})
	assert.Equal(t, []direction{}, res)
}
func TestDirReduc_Komplex01_Erfolg(t *testing.T) {
	res := DirReduc([]direction{North, South, South, East, West, North, West})
	assert.Equal(t, []direction{West}, res)
}
