package main

import "fmt"

func main() {
	ch := make(chan string, 2) // ________________
	ch <- "str1"               // |      |       |
	ch <- "str2"               // | str1 | str2  |
	fmt.Println(<-ch)          // |______|_______|
	fmt.Println(<-ch)
}
