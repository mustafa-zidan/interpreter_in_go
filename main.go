package main

import (
	"fmt"
	"os"
	"os/user"

	"moose/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Moose programming language\n", user.Username)
	fmt.Println("Happy Programming")
	repl.Start(os.Stdin, os.Stdout)
}
