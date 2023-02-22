package main

import "fmt"

type base struct {
	num int64
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
	// Inserting One struct to another is called embedding of structs
	base
	str string
}

func main() {
	s1 := container{base: base{
		num: 1,
	},
		str: "Hello World"}

	fmt.Println(s1)
	fmt.Printf("s1={num: %v, str: %v}\n", s1.num, s1.str)
	fmt.Println(s1.base)
	fmt.Println(s1.num)
	fmt.Println(s1.base.num)
	fmt.Println(s1.str)
}
