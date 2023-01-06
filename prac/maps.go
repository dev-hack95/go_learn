package main

import "fmt"

func main() {
	result := map[string]int{
		// Key : value
		"English": 89,
		"Maths":   99,
		"History": 84,
		"Science": 81,
	}
	fmt.Println(result)
	fmt.Println(result["Maths"])
	result["Sanskrit"] = 81
	fmt.Println(result)
	delete(result, "Sanskrit")
	fmt.Println(result)

	//m := map[[3]int]string{}
	//fmt.Println(m)
}
