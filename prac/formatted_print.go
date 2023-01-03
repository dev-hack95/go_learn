package main

import "fmt"

func main() {
	fmt.Printf("%-15v $%3v\n", "SpaceX", 94)
	var i int
	i = 17
	fmt.Println(i)
	// Value assigned to var varible can be mutable or can be changed
	i = 10
	fmt.Println(i)
	const j = 11
	// Value assigned to const datatype is immutable or can't be changed
	fmt.Println(j)

	// var k = 20
	// var m = 30
	// Insted of declaring two var varibles we can declare k and m in single var
	// Same goes with constant
	var (
		k = 20
		m = 30
	)
	fmt.Println(k)
	fmt.Println(m)
}
