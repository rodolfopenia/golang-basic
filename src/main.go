package main

import "fmt"

func main() {
	module := 8 % 2

	switch module {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("Es impar")
	}

	value := 200
	switch {
	case value > 100:
		fmt.Println("Es mayor a 100")
	case value < 0:
		fmt.Println("Es menor de 0")
	default:
		fmt.Println("No condición")
	}
}
