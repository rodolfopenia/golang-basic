package main

import "fmt"

func evaluateNumber(number int) string {
	var result string
	if number%2 == 0 {
		result = "par"
	} else {
		result = "impar"
	}
	return result
}

func login(user, password string) bool {
	var result bool
	if user == "user" && password == "123" {
		result = true
	} else {
		result = false
	}
	return result
}

func main() {
	const number int = 7
	result := evaluateNumber(number)
	fmt.Printf("El número %d es %s\n", number, result)

	const user string = "user"
	const password string = "1234"
	login := login(user, password)
	if login {
		fmt.Println("Login exitoso")
	} else {
		fmt.Println("Error al realizar login")
	}
}
