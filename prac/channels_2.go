package main

import "fmt"

func sums(s []int, c chan int) {
	sum := 0
	for _, i := range s { // s is slice that conatin integer data-type
		sum += i
	}
	// Send sum to chan c
	c <- sum
}

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8}

	c := make(chan int)
	go sums(s[:len(s)/2], c)
	go sums(s[len(s)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(x + y)
}
