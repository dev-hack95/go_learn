package main

import "fmt"

type data struct {
	Name   string
	Age    int
	Result map[string]int
}

func area() {
	data_1 := data{
		Name: "abc",
		Age:  21,
		Result: map[string]int{
			"Maths":   91,
			"Science": 81,
			"English": 89,
		},
	}
	for k, v := range data_1.Result {
		fmt.Printf("%v %v\n", k, v)
	}
}

func main() {
	area()
}
