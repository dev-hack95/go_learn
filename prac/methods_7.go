// Methods and pointer indirection
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	x, y float64
}

// Method
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

// Function
func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func main() {
	v1 := Vertex{x: 3, y: 4}
	fmt.Println(v1.Abs())
	fmt.Println(AbsFunc(v1))

	p := &Vertex{x: 3, y: -4}
	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p)) // Pointer But Isolated
}
