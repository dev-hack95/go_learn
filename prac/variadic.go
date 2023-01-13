package main

import "fmt"

func sum(nums ...int) {
	// A function that is called with the varying numbers is known as varidic function
	// Variadic functions can be called with any number of trailing arguments
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	// Here we are calling the function with different
	sum(1, 2, 3)
	sum(1, 2, 3, 4, 5, 5)

	nums := []int{1, 2, 3, 4, 5, 6, 7}
	sum(nums...)

}
