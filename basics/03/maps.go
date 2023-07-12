package main

import (
	"fmt"
	"math"
	"strings"
)

type Vertex2 struct {
	Lat, Long float64
}

var m map[string]Vertex2

func maps() {
	fmt.Println()
	fmt.Println("-- maps --")

	if m == nil {
		fmt.Println("this map is nil !")
	}

	m = make(map[string]Vertex2)

	m["bell labs"] = Vertex2{40.68433, -74.39967}
	fmt.Printf("m: %v\n", m)

	m2 := make(map[string]int)
	fmt.Printf("m2: %v\n", m2)

	m2["answer"] = 42
	fmt.Printf("m2: %v\n", m2)

	m2["answer"] = 93
	fmt.Printf("m2: %v\n", m2)

	delete(m2, "answer")
	fmt.Printf("m2: %v\n", m2)

	v, ok := m2["answer"]
	fmt.Printf("value: %d - present: %t\n", v, ok)

	// function values
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compoute(hypot))
	fmt.Println(compoute(math.Pow))

	pos, neg := adder(), adder()

	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}

	fmt.Println("-- fibonacci --")

	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

func WordCount(s string) map[string]int {
	result := make(map[string]int)

	for _, v := range strings.Fields(s) {
		result[v]++
	}

	return result
}

func compoute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func fibonacci() func() int {
	prev := 1
	num := 0
	return func() int {
		defer func() {
			temp := prev + num
			prev = num
			num = temp
		}()

		return num
	}
}
