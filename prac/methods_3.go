package main

import (
	"fmt"
	"math"
)

// You can declare a method on non-struct types
// empty struct type or non-struct type

type MyFloat float64

/* You can only declare a method with a receiver whose type is defined in the same package as the method.
   You cannot declare a method with a receiver whose type is defined in another package (which includes the built-in types such as int) */

func (f MyFloat) test() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt(3))
	fmt.Println(f.test())
}
