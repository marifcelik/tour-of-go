package main

import (
	"fmt"
)

func fibonacci_single(c, quit chan int) {
	x, y := 0, 1
	var done bool
	for !done {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			done = true
		}
	}
}

func select_() {
	fmt.Println("\n-- select --")

	c := make(chan int)
	q := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		q <- 1
	}()
	fibonacci_single(c, q)

	// default selection
	fmt.Println("\n-- default selection --")
	fmt.Println("/* commented */")

	// i commented this line because i cant wait it in development
	/*
		tick := time.Tick(500 * time.Millisecond)
		boom := time.After(1 * time.Second)

		for {
			select {
			case <-tick:
				fmt.Printf("tick ")
			case <-boom:
				fmt.Println("\nbooom !")
				return
			default:
				fmt.Printf(".")
				time.Sleep(50 * time.Millisecond)
			}
		}
	*/
}
