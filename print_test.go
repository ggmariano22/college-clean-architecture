package main

import (
	"fmt"
	"regexp"
	"testing"
)

func TestPrintHello(t *testing.T) {
	name := "College"
	text := fmt.Sprintf("Hello, %s", name)

	want := regexp.MustCompile(text)
	msg := fmt.Sprintf("Hello, %s", name)

	if !want.MatchString(msg) {
		t.Fatalf(`Hello("Gladys") = %q, want match for %#q, nil`, msg, want)
	}
}
