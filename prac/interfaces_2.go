package main

import "fmt"

type speech interface {
	Speak() string
}

// Advantage no 2 we can use Speak() func n number of times

type dog struct{}

func (d dog) Speak() string {
	return "Bhaw"
}

type cat struct{}

func (c cat) Speak() string {
	return "Mew"
}

func main() {
	rs := []speech{dog{}, cat{}}
	for _, r := range rs {
		fmt.Println(r.Speak())
	}
}
