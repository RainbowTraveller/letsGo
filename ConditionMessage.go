package main

import "fmt"

func main() {

	mesaageLen := 10
	maxMessageLen := 20

	fmt.Println("Trying to send message with length", mesaageLen)

	if mesaageLen > maxMessageLen {
		fmt.Println("Error: Message length exceeds maximum allowed length.")
	} else {
		fmt.Println("Message sent successfully.")
	}
}
