// Runes
/* A character is defined using “code points” in Unicode.
   Go language introduced a new term for this code point called rune. */
package main

import "fmt"

func main() {

	const s = 'a'

	s_rune := rune(s)
	// Prints ASCII value of a
	fmt.Println(s_rune) // 97

	test := "Hello World"
	test_rune := []rune(test)
	fmt.Println(test_rune) // [72 101 108 108 111 32 87 111 114 108 100]

}
