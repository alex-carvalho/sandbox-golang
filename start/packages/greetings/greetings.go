package greetings

import (
	"fmt"

	"example.com/greetings/subpackage"
)

// Hello returns a greeting for the named person.
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	subpackage.PrintDate()
	return message
}

func privateFunction() {
	fmt.Println("This is a private function")
}
