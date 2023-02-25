package main

import "fmt"

func main() {
	for i := 4; i <= 10; i++ {
		defer fmt.Println(i)
	}

	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
}
