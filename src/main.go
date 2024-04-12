package main

import "fmt"

func main() {
	// Declare constants
	const pi float64 = 3.14
	const pi2 = 3.1415
	fmt.Println("pi", pi)
	fmt.Println("pi2", pi2)

	// Declare int variables
	base := 12
	var height int = 14
	var area int
	fmt.Println("Base", base)
	fmt.Println("Altura", height)
	fmt.Println("Área", area)

	// Zero values
	var a int     // print 0
	var b float64 // print 0
	var c string  // print empty
	var d bool    // print false

	fmt.Println(a, b, c, d)

	const squareBase = 10
	squareArea := squareBase * squareBase
	fmt.Println("Area cuadrado:", squareArea)
}
