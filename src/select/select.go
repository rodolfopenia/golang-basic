package select1

import "fmt"

func message(text string, c chan string) {
	c <- text
}

func select1() {
	c := make(chan string, 2)
	c <- "Mensaje 1"
	c <- "Mensaje 2"
	fmt.Println(len(c), cap(c))

	// Range y Close
	close(c)
	// Una vez que se cierra no se puede agregar más mensajes
	//c <- "Mensaje 3"
	for message := range c {
		fmt.Println(message)
	}

	// select
	email1 := make(chan string)
	email2 := make(chan string)

	go message("Mensaje1", email1)
	go message("Mensaje2", email2)

	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email1:
			fmt.Println("Email recibido de email1", m1)
		case m2 := <-email2:
			fmt.Println("Email recibido de email2", m2)
		}
	}
}
