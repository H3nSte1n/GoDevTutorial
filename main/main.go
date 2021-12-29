package main

import (
	"fmt"
	"log"
	"example.com/greetings"
)

func main() {
	log.SetPrefix("[greetings]: ")
	log.SetFlags(0)

	names := []string{"Alice", "Bob", "Charlie"}

	messages, err := greetings.RandomGreetings(names)

	if(err != nil) {
		log.Fatalf("%s", err)
	}

	fmt.Println(messages)
}