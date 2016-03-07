package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func walkHelper(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		walkHelper(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		walkHelper(t.Right, ch)
	}
}
func Walk(t *tree.Tree, ch chan int) {
	walkHelper(t, ch)
	close(ch)
}

//func Walk(t *tree.Tree, ch chan int) {
//	if t.Left != nil {
//		Walk(t.Left, ch)
//	}
//	ch <- t.Value
//	if t.Right != nil {
//		Walk(t.Right, ch)
//	}
//}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	c1 := make(chan int)
	c2 := make(chan int)
	go Walk(t1, c1)
	go Walk(t2, c2)
	for i := range c1 {
		if i != <-c2 {
			return false
		}
	}
	return true
}
func main() {
	t := tree.New(1)
	c := make(chan int, 10)
	go Walk(t, c)
	//i := 0
	//for i < cap(c) {
	//	fmt.Println(<-c)
	//	i++
	//}
		for i := range c {
			if c != nil {
	fmt.Println(i)
			}
		}
	fmt.Println(Same(tree.New(1), tree.New(1)))
}

