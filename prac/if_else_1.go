package main

import "fmt"

func main() {
	person := map[string]string{
		"person_1": "Rushi",
		"person_2": "abcd",
		"person_3": "efgh",
	}
	if pop, ok := person["person_1"]; ok {
		fmt.Println(pop)
	}
}
