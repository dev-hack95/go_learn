package main

import "fmt"

type I interface {
	test()
}

type r struct {
	S string
}

func (r1 *r) test() {
	fmt.Println(r1.S)
}

func main() {
	i := &r{"Hello"}
	i.test()
}
