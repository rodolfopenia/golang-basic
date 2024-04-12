package main

import "fmt"

func main() {

	const squareBase = 10
	squareArea := squareBase * squareBase
	fmt.Println("Area cuadrado:", squareArea)

	x := 10
	y := 50

	// Add
	result := x + y
	fmt.Println("Suma:", result)

	// Rest
	result = y - x
	fmt.Println("Resta:", result)

	// Multiply
	result = x * y
	fmt.Println("Multiplicar", result)

	// Split
	result = y / x
	fmt.Println("División: ", result)

	// Residue
	result = y % x
	fmt.Println("Residuo:", result)

	// Increment
	x++
	fmt.Println("Incremental:", x)

	// Decrement
	x--
	fmt.Println("Decremental", x)

	// Homework
	// Find the area of rectangle, trapeze and circle.
	const base float64 = 10
	const height float64 = 5
	const minorBase float64 = 6
	const majorBase float64 = 8
	const PI float64 = 3.14
	const diameter float64 = 6

	rectangleArea := base * height
	trapezeArea := ((majorBase + minorBase) / 2) * height
	circleArea := diameter * PI

	fmt.Println("Área del rectángulo:", rectangleArea)
	fmt.Println("Área del trapecio:", trapezeArea)
	fmt.Println("Área del círculo:", circleArea)
}
