package main

import "fmt"

// Channels are the pipes that connect concurrent goroutines.
// You can send values into channels from one goroutine and receive those values into another goroutine.
func main() {
	// Create a new channel with make(chan val-type). Channels are typed by the values they convey.
	messages := make(chan string)

	go func() {
		messages <- "Hello World"
		//channel <- value-type
		//chan <- string
	}()

	msg := <-messages
	fmt.Println(msg)
}
