package main

import (
	"fmt"
	"log"

	"github.com/davionchai/learning-go/go_tutorials/t02_create_a_module/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// A slice of names.
	names := []string{"Gladys", "Samantha", "Darrin"}

	// // Request a greeting message.
	// message, err := greetings.Hello("Davion")
	// message, err := greetings.Hello("")

	// Request greeting messages for the names.
	messages, err := greetings.Hellos(names)

	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// // If no error was returned, print the returned message
	// // to the console.
	// fmt.Println(message)

	// If no error was returned, print the returned map of
	// messages to the console.
	fmt.Println(messages)
}
