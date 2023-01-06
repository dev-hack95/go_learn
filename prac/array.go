package main

import "fmt"

func main() {
	// Can only hold 3 elemets
	arr := [3]int{1, 2, 3}
	fmt.Printf("%v\n", arr)

	// Can store n numbers of elements
	arr_1 := []int{4, 5, 6, 7, 8}
	fmt.Printf("%v\n", arr_1)

	// An empty array
	var arr_2 []int
	fmt.Printf("%v\n", arr_2)
}
