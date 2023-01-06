package main

import "fmt"

func main() {
	result := make(map[string]int)
	result = map[string]int{
		// Key : value
		"English": 89,
		"Maths":   99,
		"History": 84,
		"Science": 81,
	}
	fmt.Println(result)

	// Checking if the keyword is present in map or not
	pop, ok := result["Math"]
	fmt.Println(pop, ok)

	pop_1, ok := result["Maths"]
	fmt.Println(pop_1, ok)
	//          Value, true/false
}
