package main

import (
	"fmt"
	"strings"
)

func isPalindromo(text string) {
	var textReverse string
	text = strings.ToLower(text)

	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}

	if text == textReverse {
		fmt.Println("Es palíndromo")
	} else {
		fmt.Println("No es palíndromo")
	}
}

func main() {

	slice := []string{"hola", "que", "hace"}

	for _, valor := range slice {
		fmt.Println(valor)
	}

	isPalindromo("Ama")
}
