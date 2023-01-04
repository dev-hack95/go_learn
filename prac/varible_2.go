package main

import (
	"fmt"
	"strconv"
)

func main() {

	// Converting int to float
	var i int
	i = 33
	var j float32
	j = float32(i)
	fmt.Printf("%v %T\n", j, j)

	// Converting int to string
	var k int
	k = 42
	fmt.Printf("%v %T\n", k, k)
	var l string
	l = strconv.Itoa(k)
	fmt.Printf("%v %T\n", l, l)

}
