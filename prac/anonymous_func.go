package main

import "fmt"

var test func() = func() {

}

func main() {
	test_1 := func() {
		fmt.Println("Hello anonymous")
	}
	test_1()

	test_2 := func() {
		fmt.Println("Hello anonymous func")
	}
	test_2()
}
