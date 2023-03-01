package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	// This is the function we’ll run in a goroutine. The done channel will be used to notify another goroutine that this function’s work is done.

	fmt.Println("Working...")
	time.Sleep(5 * time.Second)
	fmt.Println("done")

	done <- true
	// Send a value to notify that we’re done
}

func main() {
	done := make(chan bool, 1)
	go worker(done) // Start a worker goroutine, giving it the channel to notify on.
	<-done

	// Block until we receive a notification from the worker on the channel.
	// If you removed the <- done line from this program, the program would exit before the worker even started
}
