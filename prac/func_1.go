package main

import "fmt"

func main() {
	test := func() {
		fmt.Println("Hello world")
	}
	test()

	test_1 := func(x int) {
		fmt.Println(x)
	}
	test_1(5)

	test_2 := func(x int) int {
		return x * 2
	}
	test_2(5)
}
