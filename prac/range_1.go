package main

import "fmt"

type data struct {
	Name   string
	Age    int
	Result map[string]int
}

func main() {
	data_1 := data{
		Name: "abc",
		Age:  21,
		Result: map[string]int{
			"Maths":   91,
			"English": 89,
			"Science": 81,
		},
	}
	fmt.Println(data_1)
	for k, v := range data_1.Result {
		if k == "Maths" {
			fmt.Println("marks: ", v)
		}
	}
}
