package katas

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcBounces_Falscher_Param_bounce_Fehler(t *testing.T) {
	res := CalcBounces(3, 1, 1.5)
	assert.Equal(t, -1, res)
}
func TestCalcBounces_Falscher_Param_h_Fehler(t *testing.T) {
	res := CalcBounces(0, 0.5, 1.5)
	assert.Equal(t, -1, res)
}
func TestCalcBounces_Falscher_Param_window_Fehler(t *testing.T) {
	res := CalcBounces(5, 0.5, 5)
	assert.Equal(t, -1, res)
}
func TestCalcBounces_3x_Erfolg(t *testing.T) {
	res := CalcBounces(3, 0.66, 1.5)
	assert.Equal(t, 3, res)
}
func TestCalcBounces_5x_Erfolg(t *testing.T) {
	res := CalcBounces(4, 0.66, 1.5)
	assert.Equal(t, 5, res)
}
