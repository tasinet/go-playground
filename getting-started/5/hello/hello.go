package main

import (
	"example.com/greetings"
	"log"
	"strconv"
)

// init function runs on initialisation
// func init() {
//   fmt.Println("init ran")
// }

// main function runs when it is a main package
func main() {
	log.SetPrefix("greetings: ")
	// log.SetFlags(0)

	// messages, err := greetings.Hellos([]string{}) // to see error

	names := []string{"Tasos", "Stefania", ""}
	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("messages:", messages)

	for name, greeting := range messages {
		log.Println("Greeting for", strconv.Quote(name), "is", strconv.Quote(greeting))
	}
}
