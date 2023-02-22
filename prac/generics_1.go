package main

import "fmt"

func newgenericfunc[age int64 | float32](myage age) {
	fmt.Println(myage)
}

func main() {
	var testage int64 = 43
	var testage1 float32 = 43.5

	newgenericfunc(testage)
	newgenericfunc(testage1)

}
