package main

import "fmt"

type List[T any] struct {
	value T
	next  *List[T]
	head  *List[T]
	tail  *List[T]
}

func (l *List[T]) push(v T) {
	node := &List[T]{value: v}
	node.next = nil
	if l.tail == nil {
		l.head = node
		l.tail = node
	} else {
		l.tail.next = node
		l.tail = node
	}
}

func (l List[T]) print() {
	fmt.Printf("List: [")

	for i := l.head; i != nil; i = i.next {
		if i == l.tail {
			fmt.Printf("%v", i.value)
		} else {
			fmt.Printf("%v, ", i.value)
		}
	}
	fmt.Printf("]\n")
}
