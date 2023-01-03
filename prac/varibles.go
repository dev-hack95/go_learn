package main

import (
	"fmt"
)

var i int = 3

const j = 4

// var i int
// i = 3
// This is not valid if you are assigning any value to varible then you should assigned it in main
// Code on line 5 works beacuse you are assigning the value at the same time of decleration

func main() {
	//i = 2
	fmt.Println(i)
	fmt.Println(j)

	// Strings

	var str string
	var str1 string
	var str2 string

	str1 = "hello"
	str2 = " world"
	str = str1 + str2
	fmt.Printf(str)

}
