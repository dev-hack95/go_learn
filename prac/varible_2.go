package main

import "fmt"

func main() {

	// Converting int to float
	var i int
	i = 33
	var j float32
	j = float32(i)
	fmt.Printf("%v %T", j, j)

}
