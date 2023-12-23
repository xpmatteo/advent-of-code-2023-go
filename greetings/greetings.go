package greetings

import "fmt"

// Hello returns a greeting for the named person.
func Hello(name string) string {
	if name == "" {
		name = "world"
	}
	return fmt.Sprintf("Hello, %v", name)
}
