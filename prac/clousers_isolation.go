package main

import "fmt"

func newCounter() func() int {
	n := 0
	return func() int {
		n++
		return n
	}
}

func main() {
	counter := newCounter()
	fmt.Println(counter())
	fmt.Println(counter())
}
