package main

import "fmt"

func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

func main() {
	arr := []int{3, 5, 7, 9}
	fmt.Println(Index(arr, 5))

	arrS := []string{"hi", "my", "name", "is", "arif"}
	fmt.Println(Index(arrS, "furkan"))

	myList := List[int]{}
	myList.push(3)
	myList.push(5)
	myList.push(7)
	myList.print()

	myList2 := List[string]{}
	myList2.push("deneme")
	myList2.push("denem 2")
	myList2.print()
}
