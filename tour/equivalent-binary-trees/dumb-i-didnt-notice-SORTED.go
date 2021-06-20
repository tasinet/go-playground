package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func WalkChild(t *tree.Tree, ch chan int) {
	ch <- t.Value
	if t.Left != nil {
		WalkChild(t.Left, ch)
	}
	if (t.Right != nil) {
		WalkChild(t.Right, ch)
	}
}

func Walk(t *tree.Tree, ch chan int) {
	WalkChild(t, ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	var values []int
	for v1 := range ch1 {
		values = append(values, v1)
	}
	for v2 := range ch2 {
		var found bool
		values, found = Remove(values, v2)
		if !found {
			return false
		}
	}
	if len(values) != 0 {
		return false
	}
	return true
}

func Remove(list []int, value int) ([]int, bool) {
	for i, v := range list {
		if v == value {
			list[i] = list[len(list)-1]
			return list[:len(list)-1], true
		}
	}
	return list, false
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
	fmt.Println(Same(tree.New(2), tree.New(2)))
}


