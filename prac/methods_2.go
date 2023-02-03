package main

import "fmt"

/* Go automatically handles conversion between values and pointers for method calls.
   Use a pointer receiver type to avoid copying on method calls
   or to allow the method to mutate the receiving struct. */

type rect struct {
	width, height float64
}

type squr struct {
	length float64
}

func (r *rect) area() float64 {
	return r.width * r.height
}

func (s squr) area() float64 {
	return s.length * s.length
}

func main() {
	r1 := rect{width: 4, height: 5}
	fmt.Println(r1.area())

	s1 := squr{length: 5}
	fmt.Println(s1.area())
}
