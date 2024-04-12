package main

import "fmt"

func main() {
	helloMessage := "Hello"
	worldMessage := "World"

	// printLn - Add line break
	fmt.Println(helloMessage, worldMessage)
	fmt.Println(helloMessage, worldMessage)

	// printF
	nombre := "Platzi"
	cursos := 500
	fmt.Printf("%s tiene más de %d cursos\n", nombre, cursos)
	fmt.Printf("%v tiene más de %v cursos\n", nombre, cursos)

	// Sprintf
	message := fmt.Sprintf("%s tiene más de %d cursos\n", nombre, cursos)
	fmt.Println(message)

	fmt.Printf("helloMessage: %T\n", helloMessage)
	fmt.Printf("Cursos: %T\n", cursos)
}
