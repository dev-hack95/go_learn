package main

import "fmt"

type person struct {
	name  string
	age   int
	city  string
	state string
	// Struct is a collection type
}

func main() {
	person_1 := person{
		name:  "Rushi",
		age:   21,
		city:  "Nanded",
		state: "Maharashtra",
	}
	person_2 := person{
		name:  "abc",
		age:   32,
		city:  "Nagpur",
		state: "Maharashtra",
	}

	fmt.Println(person_1)
	fmt.Println(person_2)
}
