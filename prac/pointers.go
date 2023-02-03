package main

import "fmt"

func test_1(ival int) {
	ival = 0
}

func test_2(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("Starting Value: ", i)

	test_1(i)
	fmt.Println("Test_1 Value: ", i)

	test_2(&i)
	fmt.Println("Test_2 Value: ", i)

	fmt.Println("Pointer Value: ", &i)
}
