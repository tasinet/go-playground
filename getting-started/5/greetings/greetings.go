package greetings

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Hellos(names []string) (map[string]string, error) {
	if len(names) == 0 {
		return nil, errors.New("no names provided")
	}

	messages := make(map[string]string)

	for idx, name := range names {
		message, err := Hello(name)
		if err != nil {
			log.Println("Error in index:", idx, "error:", err.Error())
			continue
		}
		messages[name] = message
	}
	return messages, nil
}

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("no name")
	}

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func randomFormat() string {
	formats := []string{
		"Hello %v!",
		"Hi %v!",
		"¡Hola %v!",
	}
	return formats[rand.Intn(len(formats))]
}
