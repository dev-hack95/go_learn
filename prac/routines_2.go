package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("Hello")
	go f("Hello World")

	go func(msg string) {
		fmt.Println(msg)
	}("going")
	time.Sleep(3 * time.Second)
	fmt.Println("done")
}
