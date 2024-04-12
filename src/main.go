package main

import "fmt"

func main() {
	// Defer - execute to end
	defer fmt.Println("Hola")
	fmt.Println("Mundo")

	// Continue and Break
	for i := 0; i < 10; i++ {
		fmt.Println(i)

		if i == 2 {
			fmt.Println("Es 2")
			continue
		}

		if i == 8 {
			fmt.Println("Break")
			break
		}
	}
}
