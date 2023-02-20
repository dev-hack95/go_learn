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

}
