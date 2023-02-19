package main

import "fmt"

type I interface {
	test()
}

type T struct {
	S string
}

func (t *T) test() {
	if t == nil {
		fmt.Println("<nil>")
		return
	} else {
		fmt.Println(t.S)
	}
}

func main() {
	var i I

	var t *T
	i = t
	i.test()

	i = &T{"Hello"}
	i.test()

}
