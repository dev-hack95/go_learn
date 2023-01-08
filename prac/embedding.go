package main

import "fmt"

type Animal struct {
	Name   string
	Origin string
}

type Bird struct {
	// Inserting One struct to another is called embedding of structs
	Animal
	Speed  float32
	Canfly bool
}

func main() {
	b := Bird{}
	b.Name = "Emu"
	b.Origin = "Austrila"
	b.Speed = 48.5
	b.Canfly = true

	fmt.Println(b)
}
