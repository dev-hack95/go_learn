package main

import "fmt"

func main() {
	var i bool
	i = true
	fmt.Printf("%v %T\n", i, i)

	n := 1
	m := 1 == 2
	fmt.Printf("%v %T\n", n, n)
	fmt.Printf("%v %T\n", m, m)

	if n == 1 {
		fmt.Printf("True\n")
	} else {
		fmt.Printf("False")
	}
}
