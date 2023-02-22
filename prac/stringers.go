package main

import (
	"fmt"
	"reflect"
)

/*
type Stringer interface {
	String() string
}

A Stringer is a type that can describe itself as a string. The fmt package (and many others) look for this interface to print values.

*/

type person struct {
	Name string
	Age  int
}

func (p person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	s1 := person{"Arthur Dent", 42}
	fmt.Println(s1)
	fmt.Println(reflect.ValueOf(s1.Name).Kind())
	fmt.Println(reflect.ValueOf(s1.Age).Kind())
}
