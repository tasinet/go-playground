package main

import (
	"example.com/greetings"
	"fmt"
)

func main() {
	message := greetings.Hello("Tasos")
	fmt.Println(message)
}
