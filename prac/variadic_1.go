package main

import (
    "fmt"
)

func test(strs ...string) {
    for _, str := range strs {
        fmt.Println(str)
    }
}

func main() {
    strs := []string{"One", "Two", "Three"}
    test(strs...)
}
