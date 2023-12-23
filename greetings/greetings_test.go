package greetings

import (
	"regexp"
	"testing"
)

func TestHelloEmpty(t *testing.T) {
	actual := Hello("")
	want := "Hello, world"
	if actual != want {
		t.Fatalf(`Hello("") = %q, want %v`, actual, want)
	}
}

func TestHelloName(t *testing.T) {
	name := "Pippoz"
	want := regexp.MustCompile(`\b` + name + `\b`)
	actual := Hello(name)
	if !want.MatchString(actual) {
		t.Fatalf(`Hello("%v") = %q, want match for %#q`, name, actual, want)
	}
}
