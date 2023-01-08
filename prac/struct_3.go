package main

import "fmt"

type person struct {
	Name string
	Age  int
}

func test() {
	person_1 := person{
		Name: "abc",
		Age:  21,
	}
	fmt.Println(person_1)
}

func main() {
	test()

	// Anonymous struct
	aDoctor := struct{ name string }{name: "Abc"}
	fmt.Println(aDoctor)
}
