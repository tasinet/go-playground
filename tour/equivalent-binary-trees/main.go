package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)

// Walk recursive
func WalkChild(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		WalkChild(t.Left, ch)
	}
	ch <- t.Value
	if (t.Right != nil) {
		WalkChild(t.Right, ch)
	}
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	WalkChild(t, ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	
	for v1 := range ch1 {
		v2, ch2_open := <- ch2
		if !ch2_open {
			// channel 2 ran out but channel 1 had stuff
			return false
		}
		if v1 != v2 {
			return false
		}
	}
	_, ch2_open := <- ch2
	if ch2_open {
		// channel2 has more stuff
		return false
	}
	return true
}


func main() {
	fmt.Println("true?", Same(tree.New(1), tree.New(1)))
	fmt.Println("false?", Same(tree.New(1), tree.New(2)))
	fmt.Println("true?", Same(tree.New(2), tree.New(2)))
	fmt.Println("false?", Same(tree.New(2), tree.New(3)))
}
