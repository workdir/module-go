package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")

	names := []string{
		"Chris Collins",
		"Donald Trump",
		"Tom Marino",
		"Bob Corker",
		"Elon Musk",
	}

	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	for _, message := range messages {
		fmt.Println(message)
	}
}
