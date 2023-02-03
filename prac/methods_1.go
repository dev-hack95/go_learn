package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	x, y float64
}

/* Go does not have classes. However, you can define methods on types.
   A method is a function with a special receiver argument.
   The receiver appears in its own argument list between the func keyword and the method name.
   In this example, the Abs method has a receiver of type Vertex named v.
   Methods can be defined for either pointer or value receiver types. Here’s an example of a value receiver. */

//  Go supports methods defined on struct types.

func (v Vertex) Abs() float64 {
	return float64(math.Sqrt(float64(v.x)*float64(v.x) + float64(v.y)*float64(v.y)))
}

func main() {
	v1 := Vertex{3, 4}
	fmt.Println(v1.Abs())

	v2 := Vertex{3, -4}
	fmt.Println(v2.Abs())

	v3 := Vertex{x: 12, y: 9}
	fmt.Println(v3)
	fmt.Println(v3.Abs())
}
