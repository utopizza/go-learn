package main

import (
	"fmt"
)

func foo(s []int) {
	s = append(s, 100)
}

func main() {
	arr1 := [3]int{0, 1, 2}
	fmt.Println(arr1)

	slice1 := arr1[:2]
	foo(slice1)

	fmt.Println(arr1)
	fmt.Println(slice1)
}
