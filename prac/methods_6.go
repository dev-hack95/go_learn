// Methods and pointer indirection
package main

import "fmt"

type Vertex struct {
	x, y float64
}

func (v *Vertex) Scale(f float64) {
	v.x = v.x * f
	v.y = v.y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.x = v.x * f
	v.y = v.y * f
}

func main() {
	v1 := Vertex{x: 3, y: 4} // Compare
	v1.Scale(2)              // For the statement v1.Scale(2), even though v1 is a value and not a pointer
	ScaleFunc(&v1, 10)       // Compare

	p := &Vertex{4, 3} // Compare
	p.Scale(3)         // For the statement p.Scale(3), even though p is a value and a pointer
	ScaleFunc(p, 8)    // Compare

	fmt.Println(v1, p)
}
