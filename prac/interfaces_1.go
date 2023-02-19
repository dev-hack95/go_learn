// Interfaces are basically collection of method signatures
package main

import "fmt"

type geometary interface { // In Interafces we can save the function's and we can call functions
	area() float64  // Interfaces are used in Go where polymorphism is needed.
	perim() float64 //In a function where multiple types can be passed an interface can be used

}

type rect struct {
	width, height float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func measure(g geometary) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 3, height: 4}

	measure(r)
}
