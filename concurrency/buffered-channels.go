package main

import "fmt"

func buffered_channels() {
	fmt.Println("\n-- buffered channel --")
	c := make(chan int, 2)
	c <- 1
	c <- 2
	// deadlock !
	// c <- 3
	fmt.Println(<-c)
	fmt.Println(<-c)
}
