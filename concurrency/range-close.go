package main

import "fmt"

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func range_and_close() {
	fmt.Println("\n-- range and close --")

	c := make(chan int, 12)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Printf("%v ", i)
	}
	fmt.Println()
}
