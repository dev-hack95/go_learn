package main

import "fmt"

const (
	_ = iota
	a
	b
	c
)

const (
	_ = 1
	a1
	b1
	c1
)

func main() {
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)
	fmt.Printf("%v\n", a1)
	fmt.Printf("%v\n", b1)
	fmt.Printf("%v\n", c1)
}
