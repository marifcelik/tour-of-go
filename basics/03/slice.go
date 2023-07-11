package main

import (
	"fmt"
	"reflect"
	"strings"
)

func slice() {
	q := []int{2, 3, 5, 8}
	fmt.Printf("q: %v\n", q)
	fmt.Println(reflect.TypeOf(q).Kind())
	st := []struct {
		i int
		b bool
	}{{3, true}, {7, false}}
	fmt.Printf("st: %v\n", st)

	si := []int{2, 3, 5, 7, 11, 13}
	printSlice(si)

	si = si[:0]
	printSlice(si)

	si = si[:4]
	printSlice(si)

	si = si[2:]
	printSlice(si)

	// the zero value of a slice is nil
	var nilS []int
	printSlice(nilS)

	if nilS == nil {
		fmt.Println("this slice is nil")
	}

	a := make([]int, 5)
	printSlice(a)

	b := make([]int, 0, 5)
	printSlice(b)

	b = b[:cap(b)]
	printSlice(b)

	b = b[1:]
	printSlice(b)

	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	var zeroS []int

	zeroS = append(zeroS, 2)
	printSlice(zeroS)

	zeroS = append(zeroS, 5, 8, 9, 1)
	printSlice(zeroS)

	pow := []int{1, 2, 4, 8, 16, 32, 64, 128}

	for i, v := range pow {
		fmt.Printf("2**%d=%d\n", i, v)
	}
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func Pic(dx, dy int) [][]uint8 {
	s := make([][]uint8, dy)

	for i := range s {
		s[i] = make([]uint8, dx)
		for j := range s[i] {
			s[i][j] = uint8(i*i + j*j)
		}
	}
	return s
}
