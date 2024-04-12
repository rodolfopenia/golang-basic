package main

import (
	pk "curso-golang/src/mypackage"
	"fmt"
)

type car struct {
	brand string
	year  int
}

func main() {
	var myCar pk.CarPublic
	myCar.Brand = "Ferrari"
	myCar.Year = 2020
	fmt.Println(myCar)

	pk.PrintMessage("Hola PLatzi")
}
