package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

type I interface {
	M()
}

type T struct {
	S string
}

type F float64

func interfaces() {
	fmt.Println()
	fmt.Println("-- interfaces --")

	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex2{3, 4}

	fmt.Println("a f implements Abser")
	a = f // f implenets Abser
	fmt.Println(a.Abs())
	fmt.Println("a &v implements Abser")
	a = &v // &v implements Abser
	fmt.Println(a.Abs())

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	// a = v

	t := T{"hi"}
	t.M()

	var i I
	i = &T{"whats up"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()

	var i2 I
	var t2 *T

	i2 = t2
	describe(i2)
	i2.M()

	// an interface value that holds a nil concrete value is itself non-nil.
	// so this comparison is never true
	// if i2 == nil {
	// 	fmt.Println("i2 is nil")
	// }

	i2 = &T{"hello"}
	describe(i2)
	i2.M()

	// nil interface value
	var i3 I
	describe(i3)
	// runtime error because there is no type or value in this tuple
	// i3.M()

	// empty interface
	fmt.Println()
	fmt.Println("empty interface")
	var i4 interface{}
	fmt.Printf("(%v %T)\n", i4, i4)

	i4 = 73
	fmt.Printf("(%v %T)\n", i4, i4)

	i4 = "lord voldemort"
	fmt.Printf("(%v %T)\n", i4, i4)

}

type Vertex2 struct {
	X, Y float64
}

func (v *Vertex2) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
// func (t T) M() {
// 	fmt.Println(t.S)
// }

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func (f F) M() {
	fmt.Println(f)
}

func describe(i I) {
	fmt.Printf("(%v %T)\n", i, i)
}
