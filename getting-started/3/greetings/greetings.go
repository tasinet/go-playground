package greetings

import (
  "errors"
  "fmt"
)

func Hello(name string) (string, error) {
  if name == "" {
    return "", errors.New("no name")
  }

  message := fmt.Sprintf("Hello %v. Welcome!", name)
  return message, nil
}
