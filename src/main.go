package main

import "fmt"

func main() {

	// Array - Son inmutables
	var arr [4]int
	arr[0] = 1
	arr[1] = 2
	fmt.Println(arr, len(arr), cap(arr))

	// Slice
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(slice, len(slice), cap(slice))

	// Methods slice
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[4:])

	// Append
	slice = append(slice, 7)
	fmt.Println(slice)

	// Append new slice
	newSlice := []int{8, 9, 10}
	slice = append(slice, newSlice...)
	fmt.Println(slice)
}
