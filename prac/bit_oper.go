package main

import "fmt"

func main() {
	a := 10 // 1010
	b := 3  // 0011

	fmt.Println(a & b)  // AND operator
	fmt.Println(a | b)  // OR operator
	fmt.Println(a ^ b)  // Exclusive OR operator
	fmt.Println(a &^ b) // AND-NOT operator
}
