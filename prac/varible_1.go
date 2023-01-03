package main

import "fmt"

var i int = 10

// When you Declare a varible out the package which is in smaller case then it is limited to the domain of the package

var K int = 50

// When you Declare a varible out the package which is in upper case then it is a global level varible

func main() {
	var i int
	// we can't declare the same varible twice in fuction or outside the fuction we can declare it in outside and inside for only one time
	//# command-line-arguments
	//.\varible_1.go: i redeclared in this block
	//    .\varible_1.go: other declaration of i
	//var i int = 50
	fmt.Println(i)

	// When you declare a varibale with j := 3 or something other you have to use that varible otherwise it will give an conpile error
	//j := 3
	//fmt.Println(j)

	// When you declare a value var i int like that there is no error generated beacuse it has assined a value 0
}
