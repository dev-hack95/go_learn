package main

import "fmt"

func main() {
	var i interface{} = "Hello world"
	// Checking that inteface i contain string or not 's' is a string and 'ok' is bool datatype
	// It is a contition statement if s == 'str' then print str and return true else return false
	s1, ok := i.(string)
	fmt.Println(s1)
	fmt.Println(ok)
	fmt.Println()

	// Here i is not a float therefore it will return false
	s2, ok := i.(float64)
	fmt.Println(s2) // 0
	fmt.Println(ok) // false
	fmt.Println()

	// Why it is returning zero (0) in golang if you don't assign any value to int or float datatype it will return zero

	s3, ok := i.([]byte)
	fmt.Println(s3)
	fmt.Println(ok)
}
