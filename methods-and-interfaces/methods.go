package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// where is the constructor ?
type MyFloat float64

func methdos() {
	fmt.Println("-- methods --")

	v := Vertex{6, 8}
	fmt.Println(v.Abs())
	// regular version
	fmt.Println(Absf(v))

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	fmt.Println()
	fmt.Println("before scale")
	fmt.Printf("v: %v\n", v)
	v.Scale(3)
	fmt.Println("after scale")
	fmt.Printf("v: %v\n", v)

	fmt.Println()
	fmt.Println("calling regular function")
	Scalef(&v, 2)
	fmt.Printf("v: %v\n", v)

	fmt.Println("pointer vertex")
	vp := &Vertex{5, 12}
	vp.Scale(2)
	Scalef(vp, 2)
	fmt.Println(vp)
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// regular function with same functionality
func Absf(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X *= f
	v.Y *= f
}

// regular function version
func Scalef(v *Vertex, f float64) {
	v.X *= f
	v.Y *= f
}

// i dont understand yet, where is the
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}
