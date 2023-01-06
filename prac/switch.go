package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	switch i {
	case 1:
		fmt.Printf("One\n")
	case 2:
		fmt.Printf("Two\n")
	case 3:
		fmt.Printf("Three\n")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Class at 3 pm")
	default:
		fmt.Printf("Go to college")
	}
}
