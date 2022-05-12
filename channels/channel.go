package channels

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	walk(t, ch)
	close(ch)
	// var nodeQueue []*tree.Tree = make([]*tree.Tree, 0)
	// var aktNode = t

	// for {
	// 	if aktNode.Left != nil {
	// 		nodeQueue = append(nodeQueue, aktNode.Left)
	// 	}
	// 	if aktNode.Right != nil {
	// 		nodeQueue = append(nodeQueue, aktNode.Right)
	// 	}
	// 	ch <- aktNode.Value

	// 	if len(nodeQueue) > 0 {
	// 		aktNode = nodeQueue[len(nodeQueue)-1]
	// 		nodeQueue = nodeQueue[0 : len(nodeQueue)-1]
	// 	} else {
	// 		close(ch)
	// 		break
	// 	}
	// }
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
	fmt.Println("T1: ", t1)
	fmt.Println("T2: ", t2)
	result := false
	c1 := make(chan int)
	c2 := make(chan int)
	go Walk(t1, c1)
	go Walk(t2, c2)
	t2Values := make([]int, 11)

	counter := 0
	for i := range c2 {
		t2Values[counter] = i
		counter++
	}
	fmt.Println("t2Values: ", t2Values)
	for i := range c1 {
		fmt.Println("i aus c1: ", i)
		var found bool = false
		for _, j := range t2Values {

			if i == j {
				fmt.Printf("Found: %v-%v", i, j)
				found = true
				break
			}
		}
		result = found
	}

	return result
}
