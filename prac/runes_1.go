package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "สวัสดี"

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}

	fmt.Println("Rune Count", utf8.RuneCountInString(s))
}
