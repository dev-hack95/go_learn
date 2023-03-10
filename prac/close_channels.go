package main

import "fmt"

// Note : Only the sender should close a channel, never the receiver. Sending on a closed channel will cause a panic.
// Fibbonaci is the sender
func fibbonci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

// Note: Channels aren't like files; you don't usually need to close them.
//       Closing is only necessary when the receiver must be told there are no more values coming,
//       such as to terminate a range loop.

func main() {
	c := make(chan int, 10)
	go fibbonci(cap(c), c)
	for i := range c { // The loop for i := range c receives values from the channel repeatedly until it is closed
		fmt.Println(i)
	}
}
