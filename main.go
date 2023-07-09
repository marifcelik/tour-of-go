package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"time"
)

var c, python, java bool

func main() {
	fmt.Println("hello arif")
	fmt.Println("year is", time.Now().Year())

	fmt.Println("random int", rand.Intn(10))
	fmt.Printf("now you have %g problems\n", math.Sqrt(math.Pi))

	fmt.Println("func toplamı ", add(1, 2, 3))

	str1, str2 := swap("detective", "true")
	fmt.Println(str1, str2)

	fmt.Println(split(9))

	var i int
	fmt.Println(i, c, python, java)

	var c2, python2, java2 = true, false, "no!"
	fmt.Println(c2, python2, java2)

	c3, python3, java3 := false, true, "yes!"
	fmt.Println(c3, python3, java3)

	var (
		toBe   bool       = false
		maxInt uint64     = 1<<64 - 1
		z      complex128 = cmplx.Sqrt(-5 + 12i)
	)

	fmt.Println()
	fmt.Printf("type: %T - value: %v\n", toBe, toBe)
	fmt.Printf("type: %T - value: %v\n", maxInt, maxInt)
	fmt.Printf("type: %T - value: %v\n", z, z)

	n := 42
	f := float64(n)
	u := uint(f)
	fmt.Println(n, f, u)

	fmt.Println()
	const Truth = "no"
	fmt.Println("is this world real ?")
	fmt.Println(Truth)
}

func add(x, y, z int) int {
	return x + y + z
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
