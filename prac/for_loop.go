package main

import "fmt"

func main() {
	// Method 1
	i := 1
	for i <= 5 {
		fmt.Println(i)
		i += 1
		// or i++
	}

	// Method 2
	for j := 6; j <= 10; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Printf("loop\n")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
