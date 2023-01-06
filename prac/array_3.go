package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(arr[:])
	fmt.Println(arr[3:])
	fmt.Println(arr[3:6])
	fmt.Println(arr[:6])

	arr_1 := make([]int, 5) // Array of length 5 and capacity 5
	fmt.Println(arr_1)
	fmt.Println(len(arr_1))
	fmt.Println(cap(arr_1))

	arr_2 := make([]int, 5, 100) // Array of length 5 and capacity of 100
	fmt.Println(arr_2)
	fmt.Println(len(arr_2))
	fmt.Println(cap(arr_2))

	a := []int{} // Empty array
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))
	a = append(a, 1) // Append to to an empty array
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))

	a = append(a, 2, 3, 4, 5)
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))

	a = append(a, []int{2, 3, 4, 5}...)
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))

}
