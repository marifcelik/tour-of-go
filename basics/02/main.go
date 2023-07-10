package main

import (
	"fmt"
	"math"
	"runtime"
)

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	sum2 := 1

	for sum2 < 200 {
		sum2 += sum2
	}
	fmt.Println(sum2)

	if true {
		fmt.Println("condition is true !")
	}

	fmt.Println(sqrt(9), sqrt(-4))
	fmt.Println()
	fmt.Println(pow(3, 2, 10), pow(3, 3, 20))

	fmt.Println()
	fmt.Println("sqrt implemation: sqrt 9 is ", Sqrt(9))
	fmt.Println("sqrt implemation: sqrt 21 is ", Sqrt(21))

	fmt.Println("switch")
	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("hey its linux !")
	case "darwin":
		fmt.Println("os x, its cool")
	case "windows":
		fmt.Println("trash")
	default:
		fmt.Println(os)
	}

	defer fmt.Println("main func finished")

	fmt.Println("defer stacks - lifo")
	for i := 0; i < 4; i++ {
		defer fmt.Println(i)
	}
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func Sqrt(x float64) float64 {
	z := x / 2
	const delta = 0.0000001
	i := 0
	for {
		prevZ := z
		z -= (z*z - x) / (2 * z)
		i++
		if math.Abs(prevZ-z) < delta {
			break
		}
	}
	fmt.Println("loop count", i)

	return z
}
