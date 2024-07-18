package katas

import (
	"fmt"
	"testing"
)

func TestPrintSpiral(t *testing.T) {
	res := MakeSpiral(5)
	fmt.Println(res)
}
