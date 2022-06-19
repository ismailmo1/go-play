package greetings

import "fmt"

// return greeting for named person
func Hello(name string) string {
	// return greeting with name
	message := fmt.Sprintf("hi, %v. Welcome!", name)
	return message
}
