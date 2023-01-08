package main

import "fmt"

type data struct {
	name   string
	age    string
	result map[string]int
}

func main() {
	data_1 := data{
		name: "abc",
		age:  "17",
		result: map[string]int{
			"Maths":   99,
			"English": 89,
			"Science": 81,
		},
	}

	data_2 := data{
		name: "abd",
		age:  "17",
		result: map[string]int{
			"Maths":   100,
			"English": 100,
			"Science": 100,
		},
	}

	fmt.Println(data_1)
	fmt.Println(data_2)
}
