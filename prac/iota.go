package main

import "fmt"

const (
	a = iota
	b = iota
	c = iota
	// Iota is scoped to the constant when we create a new constant value of iota is zero again
)

const (
	a2 = iota
	b2
	c2
	// When we assign value to the first constant then it will overwite value to another within same constant
)

const (
	a3 = iota
)

const (
	a4 = 1
	b4
	c4

	// When we assign value to the first constant then it will overwite the same value to another within same constant
)

func main() {
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)
	fmt.Printf("%v\n", a2)
	fmt.Printf("%v\n", b2)
	fmt.Printf("%v\n", c2)
	fmt.Printf("%v\n", a3)
	fmt.Printf("%v\n", a4)
	fmt.Printf("%v\n", b4)
	fmt.Printf("%v\n", c4)

	// Iota is assignig diffent values to a, b and c i.e 0, 1 and 2 respectively
}
