package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func channels() {
	nums := []int{7, 5, -2, -8, 3, -12}
	c := make(chan int)
	go sum(nums[:len(nums)/2], c)
	go sum(nums[len(nums)/2:], c)

	x, y := <-c, <-c
	fmt.Println(x, y)
}
