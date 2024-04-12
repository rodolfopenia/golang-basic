package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)

	m["Jose"] = 14
	m["Pepito"] = 20

	fmt.Println(m)

	// Iterar map
	for i, v := range m {
		fmt.Println(i, v)
	}

	// Found a value in maps
	value, ok := m["Jose"]
	fmt.Println(value, ok)
}
