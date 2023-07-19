package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func Walk(t *tree.Tree, c chan int) {
	if t == nil {
		return
	}

	Walk(t.Left, c)
	c <- t.Value
	Walk(t.Right, c)
}

func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for i := 0; i < 10; i++ {
		v1, v2 := <-ch1, <-ch2
		if v1 != v2 {
			return false
		}
	}

	return true
}

func equivalent_binary_tree() {
	fmt.Println("\n-- exercise : equivalent binary tree --")

	ch := make(chan int)
	t1 := tree.New(2)
	fmt.Println(t1)

	go Walk(t1, ch)

	for i := 0; i < 10; i++ {
		fmt.Printf("%v ", <-ch)
	}
	fmt.Println()

	fmt.Println(Same(tree.New(2), tree.New(2)))
	fmt.Println(Same(tree.New(2), tree.New(4)))
}
