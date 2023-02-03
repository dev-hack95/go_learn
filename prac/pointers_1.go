package main

import "fmt"

func main() {
	var iptr *int
	*iptr = 1

	fmt.Println(iptr)
	fmt.Println(&iptr)
}
