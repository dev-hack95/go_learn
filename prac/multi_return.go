package main

import "fmt"

func test() (int, int) {
	return 3, 7
}

func main() {
	a, b := test()
	fmt.Println(a)
	fmt.Println(b)

	_, c := test()
	fmt.Println(c)

	d, _ := test()
	fmt.Println(d)
}
