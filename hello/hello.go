package main

// go mod init bondarik.net/hello
// go mod edit -replace bondarik.net/greetings=../greetings

import (
	"fmt"
	"log"

	"bondarik.net/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	//A slice of names
	names := []string{"Gladys", "Gopher", "Gopher"}

	// Request a greeting message.

	message, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned message
	// to the console.
	fmt.Println(message)
}
