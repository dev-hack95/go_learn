package main

import "fmt"

func test(strs ...string) {
	for _, str := range strs {
		fmt.Println(str)
	}
}

func main() {
	test("hello", "World")

	str := []string{"The", "golang", "is", "Used", "in", "k8s"}
	test(str...)
}
