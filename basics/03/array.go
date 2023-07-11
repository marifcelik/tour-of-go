package main

import (
	"fmt"
	"reflect"
)

func array() {
	fmt.Println()
	fmt.Println("arrays")
	var a [2]string
	a[0] = "hi"
	a[1] = "rust"
	fmt.Printf("a: %v\n", a)

	primes := [6]int{1, 3, 5, 7, 9, 11}
	fmt.Printf("primes: %v\n", primes)
	fmt.Printf("primes length %v\n", len(primes))
	fmt.Printf("primes cap %v\n", cap(primes))
	fmt.Println(reflect.TypeOf(primes).Kind())

	s := primes[1:4]
	fmt.Printf("s: %v\n", s)

	words := [4]string{"red", "plane", "violet", "cars"}
	fmt.Printf("words: %v\n", words)

	sl1 := words[0:2]
	sl2 := words[1:3]
	fmt.Printf("sl1: %v\n", sl1)
	fmt.Printf("sl2: %v\n", sl2)

	sl2[0] = "trees"
	fmt.Println(sl1, sl2)
	fmt.Printf("words: %v\n", words)
}
