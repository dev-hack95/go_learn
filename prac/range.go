package main

import "fmt"

func main() {
	// Print Sum of all elements in slice
	nums := []int{1, 2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("Sum: ", sum)

	// Print the index of 3
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index", i)
		}
	}

	// Create a dictonary and print key values

	kvs := map[string]int{
		"Person-1": 21,
		"Person-2": 22,
		"Person-3": 23,
	}

	for k, v := range kvs {
		fmt.Printf("%v %v\n", k, v)
	}

	// Ascii Values
	for i, c := range "apple" {
		fmt.Println(i, c)
	}
}
