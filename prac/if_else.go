package main

import "fmt"

func main() {
	var i int
	fmt.Printf("Enter_num: ")
	fmt.Scan(&i)
	if i%2 == 0 {
		fmt.Printf("%v is even\n", i)
	} else if i%2 != 0 {
		fmt.Printf("%v is odd\n", i)
	} else {
		fmt.Printf("The number is not in range")
	}
}
