package main

import "fmt"

func main() {

	var metrics [3]string
	fmt.Printf("%v\n", metrics)

	metrics[0] = "Andrew"
	fmt.Printf("%v\n", metrics)

	metrics[1] = "Tate"
	fmt.Printf("%v\n", metrics)

	metrics[2] = "Metrics-Solver"
	fmt.Printf("%v\n", metrics)

}
