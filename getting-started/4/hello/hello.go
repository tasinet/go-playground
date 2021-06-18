package main

import (
	"example.com/greetings"
	"fmt"
	"log"
)

// init function runs on initialisation
// func init() {
//   fmt.Println("init ran")
// }

// main function runs when it is a main package
func main() {
	log.SetPrefix("greetings: ")
	// log.SetFlags(0)

	message, err := greetings.Hello("Tasos")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
