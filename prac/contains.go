package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("You find yourself in a dimly lit cavern.")
	var command = "Walk outside"
	var exit = strings.Contains(command, "outside")
	fmt.Printf("You leave the cave: %v", exit)
}
