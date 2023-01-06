package main

import (
	"fmt"
)

func main() {
	result := make(map[string]int)
	result = map[string]int{
		"first":  1,
		"Second": 2,
		"Third":  3,
	}

	fmt.Println(result)
}
