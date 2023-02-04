// Pointers and functions
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	x, y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func Scale_1(v Vertex, f float64) {
	v.x = v.x * f
	v.y = v.y * f
}

func Scale_2(v *Vertex, f float64) {
	v.x = v.x * f
	v.y = v.y * f
}

func main() {
	v1 := Vertex{x: 3, y: 4}
	fmt.Println(Abs(v1))
	Scale_1(v1, 10)
	fmt.Println(Abs(v1))
	Scale_2(&v1, 10)
	fmt.Println(Abs(v1))

}
