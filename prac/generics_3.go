package main

import "fmt"

// Index returns the index of x in s, or -1 if not found
// Function that find the index of selected data from a slice or a list
func Index[T comparable](s []T, x T) int {
	// comparable is a useful constraint that makes it possible to use the == and != operators on values of the type

	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

func main() {
	s1 := []int{5, 10, 15, 20, -10}
	fmt.Println(Index(s1, 15))
}
