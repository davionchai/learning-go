package t02_create_module

import "fmt"

// Hello returns a greeting for the named person.
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	var message string = fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
