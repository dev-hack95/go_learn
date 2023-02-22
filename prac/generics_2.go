package main

import "fmt"

func MapKeys[K comparable, V any](s []K, m map[K]V) any {
	// MapKeys takes a map of any type and returns a slice of its keys
	// This function has two type parameters - K and V
	// K has the comparable constraint, meaning that we can compare values of this type with the == and != operators
	// V has the any constraint, meaning that it’s not restricted in any way
	// 'any' is an alias for interface{}
	// V can accept any data type
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

type List[T any] struct {
	// As an example of a generic type, List is a singly-linked list with values of any type.
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.head = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) GetAll() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func main() {

	var m = map[int]string{1: "1", 2: "4", 3: "6"}
	fmt.Println("keys", MapKeys(m))

	_ = MapKeys[int, string](make([]int, [int]{lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("list:", lst.GetAll())
}
