package main

import (
	"fmt"

	"github.com/davionchai/learning-go/go_tutorials/t02_create_module"
)

func main() {
	// Get a greeting message and print it.
	message := t02_create_module.Hello("Davion")
	fmt.Println(message)
}
