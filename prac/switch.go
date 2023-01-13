package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	switch i {
	case 1, 3, 5, 7, 9:
		fmt.Printf("Odd\n")
	case 2, 4, 6, 8, 10:
		fmt.Printf("Even\n")

		switch time.Now().Weekday() {
		case time.Saturday, time.Sunday:
			fmt.Println("Class at 3 pm")
		default:
			fmt.Printf("Go to college")
		}
	}
}
