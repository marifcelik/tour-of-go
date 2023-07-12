package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

type MyFloat float64

func main() {
	v := Vertex{6, 8}
	fmt.Println(v.Abs())
	// regular version
	fmt.Println(Absf(v))

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}

func (v Vertex) Abs() float64 {
	return math.Abs(v.X*v.X + v.Y*v.Y)
}

// regular function with same functionality
func Absf(v Vertex) float64 {
	return math.Abs(v.X*v.X + v.Y*v.Y)
}

// i dot understand yet
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}
