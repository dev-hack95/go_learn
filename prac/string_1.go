package main

import (
	"fmt"
)

func main() {
	s := "Hello WOrld"
	//fmt.Printf(string(s[2]))
	//s2 := " This is Golang"
	// fmt.Printf(s + s2)

	b := []byte(s)
	fmt.Printf("%v", b)
	// [72 101 108 108 111 32 87 79 114 108 100] []uint8
	// uint --> unsigned integer

	c := []rune(s)
	fmt.Printf("%v", c)

}
