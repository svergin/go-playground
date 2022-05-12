package channels

import (
	"fmt"
	"testing"

	"golang.org/x/tour/tree"
)

func TestWalkTree(t *testing.T) {
	ch := make(chan int, 10)
	go Walk(tree.New(1), ch)
	for i := range ch {
		fmt.Println(i)
	}

}

func TestTreesAreSame(t *testing.T) {
	isSame := Same(tree.New(1), tree.New(1))
	if !isSame {
		t.Errorf("Content of Trees should be the same")
	}
}

func TestTreesAreNotSame(t *testing.T) {
	isSame := Same(tree.New(1), tree.New(10))
	if isSame {
		t.Errorf("Content of Trees should not be the same")
	}
}

func TestChannelIteration(t *testing.T) {
	c1 := make(chan int, 3)

	go Walk(tree.New(1), c1)

	for i := range c1 {
		fmt.Printf("Wert[1]: %v\n", i)
	}
	for i := range c1 {
		fmt.Printf("Wert[2]: %v\n", i)
	}
}
