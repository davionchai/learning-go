package main

import (
	"fmt"
	"os"

	"github.com/davionchai/learning-go/hello-world/hello"
	"golang.org/x/example/hello/reverse"
)

func main() {
	fmt.Println("Hello")
	fmt.Println(reverse.String("Hello"))

	if len(os.Args) > 1 {
		fmt.Println(hello.Say(os.Args[1:]))
	} else {
		fmt.Println("Hello, world")
	}
}
