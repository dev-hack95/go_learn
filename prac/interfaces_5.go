package main

import (
	"fmt"
	"math"
)

type I interface {
	test() // Value, type; if type != null you can call any type in method while running it
}

type t1 struct {
	s1 string
}

type F float64

func (t *t1) test() {
	fmt.Println(t.s1)
}

func (f F) test() {
	fmt.Println(f)
}

func main() {
	var i I

	i = &t1{"Hello"}
	i.test()

	i = F(math.Pi)
	i.test()
}
