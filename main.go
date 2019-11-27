package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/mustafa-zidan/interpreter_in_go/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Moose programming language\n", user.Username)
	fmt.Println("Feel free to start typing in command")
	repl.Start(os.Stdin, os.Stdout)
}
