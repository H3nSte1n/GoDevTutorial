package greetings

import (
	"fmt"
	"errors"
	"math/rand"
  "time"
)

func Greeting(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func RandomGreetings(names []string) (map[string]string, error){
	if len(names) == 0 {
		return nil, errors.New("empty names")
	}

	messages := make(map[string]string)

	for _, name := range names {
		if name == "" {
			return nil, errors.New("empty name")
		}
		
		message, err := Greeting(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}

	return messages, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}