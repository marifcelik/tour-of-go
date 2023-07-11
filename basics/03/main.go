package main

import "fmt"

type User struct {
	name    string
	surname string
	age     int
}

type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1}
	v3 = Vertex{}
	b  = &Vertex{3, 4}
)

func main() {
	i := 42
	p := &i
	fmt.Println(*p)

	point := 1000
	pPoint := &point
	*pPoint = *pPoint * 2

	fmt.Println(pPoint)
	fmt.Println(*pPoint)
	fmt.Println(point)

	fmt.Println()
	age := 22
	fmt.Println("age", age)
	fmt.Println("age address", &age)
	call(age)
	fmt.Println("after call", age)
	callWithP(&age)
	fmt.Println("after callWithP", age)

	user := User{name: "tipit", surname: "celik", age: 58}
	user.age = 76
	fmt.Println(user)

	v := Vertex{1, 2}
	c := &v
	c.X = 1e9
	fmt.Println(v)
	fmt.Println(v1, v2, v3, b)
	fmt.Println(*b)
}

func UNUSED(x ...interface{}) {}

func call(x int) {
	fmt.Println("call method")
	x++
	fmt.Println("call x", x)
	fmt.Println("call x adress", &x)
}

func callWithP(x *int) {
	fmt.Println("callWithP method")
	*x++
	fmt.Println("call x", *x)
	fmt.Println("call x adress", x)
}
