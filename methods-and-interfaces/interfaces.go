package main

import (
	"fmt"
	"math"
	"strings"
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

	var is interface{} = "some string"

	s := is.(string)
	fmt.Println(s)

	s, ok := is.(string)
	fmt.Println(s, ok)

	f2, ok := is.(float64)
	fmt.Println(f2, ok)

	// panic
	// f2 = is.(float64)
	// fmt.Println(f2)

	do(83)
	do("arif")
	do(false)

	// Stringers
	person1 := Person{"arif", 22}
	person2 := Person{"furkan", 12}

	fmt.Println(person1)
	fmt.Println(person2)

	fmt.Println("\nStringers exercise")
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}

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

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

type Person struct {
	name string
	age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v - %v years old", p.name, p.age)
}

type IPAddr [4]byte

func (i IPAddr) String() string {
	var str strings.Builder
	for j := 0; j < len(i)-1; j++ {
		fmt.Fprintf(&str, "%v.", int(i[j]))
	}
	fmt.Fprintf(&str, "%v", int(i[len(i)-1]))

	return str.String()
}
