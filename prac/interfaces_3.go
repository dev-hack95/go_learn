package main

import "fmt"

type Animal interface {
	Call() string
}

type cat struct {
	name string
}

type dog struct {
	name string
}

func (c *cat) Call() string {
	return fmt.Sprintf("Cat %s", c.name)
}

func (d *dog) Call() string {
	return fmt.Sprintf("Dog %s", d.name)
}

func main() {
	var animals []Animal = []Animal{
		&cat{"One"},
		&dog{"One"},
		&cat{"Two"},
		&dog{"Two"},
	}

	for _, animal := range animals {
		fmt.Println(animal.Call())
	}
}
