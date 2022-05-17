package channels

import (
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	walk(t, ch)
	close(ch)
}

func walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		walk(t.Left, ch)
	}
	ch <- t.Value

	if t.Right != nil {
		walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	// result := false
	c1 := make(chan int)
	c2 := make(chan int)
	go Walk(t1, c1)
	go Walk(t2, c2)

	for {
		v1, v1ok := <-c1
		v2, v2ok := <-c2
		if v1 != v2 {
			return false
		}
		if v1ok && !v2ok {
			return false
		}
		if !v1ok && v2ok {
			return false
		}
		if !v1ok && !v2ok {
			break
		}

	}
	return true

}
