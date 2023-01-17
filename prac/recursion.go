package main

import "fmt"

func test(n int) int {
	if n == 0 {
		return 1
	}
	return n * test(n-1)
}

func main() {
	fmt.Print(test(7))
}
