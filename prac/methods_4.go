package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	x, y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

/* Go automatically handles conversion between values and pointers for method calls.
   Use a pointer receiver type to avoid copying on method calls
   or to allow the method to mutate the receiving struct.

   Methods with pointer receivers can modify the value to which the receiver points (as Scale_2 does here).
   Since methods often need to modify their receiver, pointer receivers are more common than value receivers. */

func (v Vertex) Scale_1(f float64) {
	v.x = v.x * f
	v.y = v.y * f
}

func (v *Vertex) Scale_2(f float64) {
	v.x = v.x * f
	v.y = v.y * f
}

func main() {
	v1 := Vertex{x: 3, y: 4}
	fmt.Println(v1.Abs()) // Output: 5
	v1.Scale_1(10)
	fmt.Println(v1.Abs()) // Output: 5
	v1.Scale_2(10)
	fmt.Println(v1.Abs()) // Output: 5 * 10 = 50
}
