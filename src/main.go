package main

import "fmt"

func normalFunction(message string) {
	fmt.Println(message)
}

func tripleArgument(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

func doubleReturn(a int) (c, d int) {
	return a, a * 2
}

func main() {
	normalFunction("Hola mundo")
	tripleArgument(1, 2, "Hola")

	value := returnValue(2)
	fmt.Println("Value:", value)

	value1, _ := doubleReturn(2)
	println(value1)
}
